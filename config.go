package apirest

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type config struct {
	router    *chi.Mux
	logger    *logrus.Logger
	doc       *openapi3.T
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

func (c *config) WithOpenapi3(doc *openapi3.T) *config {
	c.doc = doc
	return c
}

func (c *config) WithUserStore(userStore UserStore) *config {
	c.userStore = userStore
	return c
}
