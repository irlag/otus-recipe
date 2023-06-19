package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"

	"otus-recipe/app/services"

	"otus-recipe/app/api"
	"otus-recipe/app/models"
)

type Middlewares struct {
	Services    *services.Services
	HttpMetrics *HttpMetrics
}

func NewMiddlewares(services *services.Services, httpMetrics *HttpMetrics) *Middlewares {
	return &Middlewares{Services: services, HttpMetrics: httpMetrics}
}

func (m *Middlewares) ContentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		next.ServeHTTP(w, r)
	})
}

func (m *Middlewares) ResponseMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			respWriter := NewResponseWriter(w)

			defer func() {
				if api.AppRoutes["metrics"].Path == r.RequestURI {
					return
				}

				startedAt := r.Context().Value(models.RequestStartedAtKey).(time.Time)
				method := mux.CurrentRoute(r).GetName()

				m.HttpMetrics.statReqDurations.With(
					prometheus.Labels{"method": method},
				).Observe(time.Since(startedAt).Seconds())

				m.HttpMetrics.statReqCount.With(
					prometheus.Labels{"method": method, "code": strconv.Itoa(respWriter.statusCode)},
				).Inc()
			}()

			next.ServeHTTP(respWriter, r)
		})
	}
}

func (m *Middlewares) StartedAtMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startedAt := time.Now()
			r = r.Clone(context.WithValue(r.Context(), models.RequestStartedAtKey, startedAt))

			next.ServeHTTP(w, r)
		})
	}
}
