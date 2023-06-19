package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"otus-recipe/app/services"

	"otus-recipe/app/metrics"

	"otus-recipe/app/api"
	recipeApi "otus-recipe/app/api/recipe"
	appProcessors "otus-recipe/app/processors"
	db "otus-recipe/app/storage/db/sqlc"

	"otus-recipe/app/config"
)

type Server struct {
	config      *config.Config
	Prometheus  *prometheus.Registry
	Logger      *zap.Logger
	Router      *mux.Router
	Services    *services.Services
	HttpMetrics *HttpMetrics
}

func New(config *config.Config) *Server {
	logger, err := NewLogger(config.Debug)
	if err != nil {
		logger.Fatal("can't initialize zap logger", zap.Error(err))
	}

	server := &Server{
		config: config,
		Logger: logger,
	}

	server.configurePrometheus()
	server.initializeMetrics()

	appMetrics := metrics.New()
	appMetrics.MustRegisterMetrics(server.Prometheus)

	store := db.NewStore()
	err = store.Open(config.DB)
	if err != nil {
		logger.Fatal("can't initialize db store", zap.Error(err))
	}

	server.Router = NewRouter()

	server.Services = services.New(logger, config)
	processors := appProcessors.NewProcessor(store, server.Services, logger, config)

	api.NewMetricsApi(server.Prometheus).HandleMethods(server.Router)
	api.NewHealthcheckApi(processors).HandleMethods(server.Router)
	recipeApi.NewRecipeApi(processors).HandleMethods(server.Router)

	return server
}

func (s *Server) Start() error {
	url := fmt.Sprintf("%s:%s", s.config.BindAddress, s.config.Port)

	s.Logger.Info(fmt.Sprintf("starting api server at %s", url))

	corsAllowOrigin := handlers.AllowedOrigins([]string{"*"})

	middlewares := NewMiddlewares(s.Services, s.HttpMetrics)

	s.Router.Use(
		middlewares.StartedAtMiddleware(),
		middlewares.ResponseMiddleware(),
	)

	return http.ListenAndServe(url,
		handlers.CORS(corsAllowOrigin)(
			middlewares.ContentTypeApplicationJsonMiddleware(
				handlers.CompressHandler(
					handlers.LoggingHandler(os.Stdout, s.Router),
				),
			),
		),
	)
}
