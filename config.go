package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type config struct {
	router    *chi.Mux
	logger    *logrus.Logger
	userStore UserStore
}

func NewConfig() *config {
	return &config{}
}

func (c *config) WithRouter(router *chi.Mux) *config {
	c.router = router
	return c
}

func (c *config) WithLogger(logger *logrus.Logger) *config {
	c.logger = logger
	return c
}

func (c *config) WithUserStore(userStore UserStore) *config {
	c.userStore = userStore
	return c
}
