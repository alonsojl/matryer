package api

import (
	"encoding/json"
	"errors"
	"time"

	"net/http"
)

type response struct {
	Error     string `json:"error"`
	Status    string `json:"status"`
	HTTPCode  int    `json:"http_code"`
	Datetime  string `json:"datetime"`
	Timestamp int64  `json:"timestamp"`
}

func (s *server) renderResponse(w http.ResponseWriter, err error) {
	var (
		customerr Error
		resp      = response{
			Status:    "fail",
			Datetime:  time.Now().Format("2006-01-02 15:04:05"),
			Timestamp: time.Now().Unix(),
		}
	)

	if !errors.As(err, &customerr) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.HTTPCode = customerr.Code
	resp.Error = customerr.Err

	s.respond(w, resp, resp.HTTPCode)
}

func (s *server) respond(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *server) decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
