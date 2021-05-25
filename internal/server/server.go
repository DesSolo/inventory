package server

import (
	"inventory/internal/server/storage"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type RestServer struct {
	token     string
	storage   storage.Storage
	clientApi *clientApi
	adminApi  *adminApi
}

func NewRestServer(token string, s storage.Storage) *RestServer {
	rs := &RestServer{
		token:   token,
		storage: s,
	}
	rs.adminApi = &adminApi{
		rs: rs,
	}
	rs.clientApi = &clientApi{
		rs: rs,
	}

	return rs
}

func (s *RestServer) loadMiddlewares(r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
}

func (s *RestServer) loadRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(DefaultContentType("application/json"))

		// client api
		r.Route("/client", func(r chi.Router) {
			r.Use(TokenAuthMiddleware(s.token))
			r.Post("/upload", s.clientApi.Upload(true))
			r.Post("/upload/force", s.clientApi.Upload(false))
			r.Get("/is_exist", s.clientApi.IsExist)
		})

		// admin api
		r.Route("/admin", func(r chi.Router) {
			r.Get("/export", s.adminApi.Export())
			r.Route("/uploads", func(r chi.Router) {
				r.Get("/", s.adminApi.Uploads())
				r.Route("/{serial}", func(r chi.Router) {
					r.Delete("/", s.adminApi.Delete())
				})
			})

		})
	})
}

func (s *RestServer) Run(address string) error {
	r := chi.NewRouter()
	s.loadMiddlewares(r)
	s.loadRoutes(r)
	return http.ListenAndServe(address, r)
}
