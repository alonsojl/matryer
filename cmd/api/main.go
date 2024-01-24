package main

import (
	"fmt"

	"net/http"
	"os"

	api "github.com/alonsojl/matryer"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var apiPort = os.Getenv("API_PORT")

	client, err := api.Setup()
	if err != nil {
		return err
	}
	defer client.Close()

	logger := api.NewLogger()
	userStore := api.NewMySQLUserStore(client, logger)
	router := chi.NewRouter()

	config := api.NewConfig().
		WithRouter(router).
		WithLogger(logger).
		WithUserStore(userStore)

	srv := api.NewServer(config)
	addr := fmt.Sprintf(":%s", apiPort)

	logger.Infof("listening and serving %s", addr)

	return http.ListenAndServe(addr, srv)
}
