package apirest

import "github.com/go-chi/chi/v5"

func (s *server) routes() {
	s.config.router.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", s.handleUsersGet())
			r.Post("/", s.handleUsersCreate())
		})
	})
}
