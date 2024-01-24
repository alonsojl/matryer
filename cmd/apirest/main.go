package main

import (
	"fmt"

	"net/http"
	"os"

	apirest "github.com/alonsojl/matryer"
	"github.com/go-chi/chi/v5"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var apiPort = os.Getenv("API_PORT")

	client, err := apirest.MySQLConnection()
	if err != nil {
		return err
	}
	defer client.Close()

	logger := apirest.NewLogger()
	userStore := apirest.NewMySQLUserStore(client, logger)
	router := chi.NewRouter()

	config := apirest.NewConfig().
		WithRouter(router).
		WithLogger(logger).
		WithUserStore(userStore)

	srv := apirest.NewServer(config)
	addr := fmt.Sprintf(":%s", apiPort)

	logger.Infof("listening and serving %s", addr)

	return http.ListenAndServe(addr, srv)
}
