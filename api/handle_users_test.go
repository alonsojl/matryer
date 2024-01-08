package api_test

import (
	"bytes"
	"encoding/json"
	"matryer/api"
	"matryer/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	"github.com/matryer/is"
)

func TestHandleUsers(t *testing.T) {
	var (
		is     = is.New(t)
		path   = "/api/v1/users/"
		logger = api.NewLogger()
		router = chi.NewRouter()
	)

	client, err := db.Setup()
	is.NoErr(err)
	defer client.Close()

	userStore := db.NewMySQLUserStore(client, logger)
	config := api.NewConfig().
		WithRouter(router).
		WithLogger(logger).
		WithUserStore(userStore)

	srv := api.NewServer(config)
	t.Run("get users", func(t *testing.T) {
		r := httptest.NewRequest(http.MethodGet, path, nil)
		w := httptest.NewRecorder()

		srv.ServeHTTP(w, r)
		is.Equal(w.Code, http.StatusOK)
	})
	t.Run("create user", func(t *testing.T) {
		var buf bytes.Buffer
		body := struct {
			Name      string `json:"name"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Phone     string `json:"phone"`
			Age       int32  `json:"age"`
		}{
			Name:      "Jorge Luis",
			FirstName: "Alonso",
			LastName:  "Hdez",
			Email:     "alonso12.dev@gmail.com",
			Phone:     "7713037204",
			Age:       25,
		}
		err := json.NewEncoder(&buf).Encode(body)
		is.NoErr(err)

		r := httptest.NewRequest(http.MethodPost, path, &buf)
		w := httptest.NewRecorder()

		srv.ServeHTTP(w, r)
		is.Equal(w.Code, http.StatusCreated)
	})
}
