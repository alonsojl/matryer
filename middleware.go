package api

import (
	"net/http"
	"os"
	"strconv"
)

func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAdmin, _ := strconv.ParseBool(os.Getenv("IsAdmin"))
		if isAdmin {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}
