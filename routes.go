package apirest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gopkg.in/yaml.v2"
)

func (s *server) routes() {
	s.config.router.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("docs/"))))
	s.config.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
			s.respond(w, s.config.doc, http.StatusOK)
		})
		r.Get("/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
			data, err := yaml.Marshal(s.config.doc)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/x-yaml")
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		})

		r.Route("/users", func(r chi.Router) {
			r.Get("/", s.handleUsersGet())
			r.Post("/", s.handleUsersCreate())
		})
	})
}
