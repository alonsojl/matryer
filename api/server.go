package api

import (
	"matryer/db"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type server struct {
	router    *chi.Mux
	logger    *logrus.Logger
	userStore db.UserStore
}

func NewServer(router *chi.Mux, logger *logrus.Logger, userStore db.UserStore) *server {
	s := &server{
		router:    router,
		logger:    logger,
		userStore: userStore,
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
