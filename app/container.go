package app

import (
	"go.uber.org/zap"

	"otus-recipe/app/config"
	appProcessors "otus-recipe/app/processors"
	"otus-recipe/app/server"
	"otus-recipe/app/services"
	"otus-recipe/app/services/clients"
	db "otus-recipe/app/storage/db/sqlc"
	"otus-recipe/app/storage/elastic"
)

type Container struct {
	Config     *config.Config
	Log        *zap.Logger
	Store      db.Store
	Processors *appProcessors.Processors
	Services   *services.Services
}

func NewContainer(cfg *config.Config) *Container {
	logger, err := server.NewLogger(cfg.Debug)
	if err != nil {
		logger.Fatal("can't initialize zap logger", zap.Error(err))
	}

	appClients := clients.New(cfg)
	srvs := services.New(appClients)

	elasticsearch, err := elastic.New(cfg)
	if err != nil {
		logger.Fatal("init elastic error", zap.Error(err))
	}

	store := db.NewStore()
	err = store.Open(cfg.DB)
	if err != nil {
		logger.Fatal("can't initialize db store", zap.Error(err))
	}

	prcs := appProcessors.NewProcessor(store, srvs, elasticsearch, logger, cfg)

	return &Container{
		Config:     cfg,
		Log:        logger,
		Store:      store,
		Processors: prcs,
		Services:   srvs,
	}
}
