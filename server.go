package apirest

import (
	"net/http"
)

type server struct {
	config *config
}

func NewServer(config *config) *server {
	s := &server{
		config: config,
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.config.router.ServeHTTP(w, r)
}
