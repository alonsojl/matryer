package apirest

import (
	"net/http"
	"time"
)

func (s *server) handleUsersGet() http.HandlerFunc {
	type response struct {
		Status    string      `json:"status"`
		HTTPCode  int         `json:"http_code"`
		Datetime  string      `json:"datetime"`
		Timestamp int64       `json:"timestamp"`
		User      []*UserData `json:"user"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.config.userStore.GetUsers()
		if err != nil {
			s.config.logger.Errorf("error getting users: %v", err)
			s.respondErr(w, err)
			return
		}

		s.respond(w, response{
			Status:    "success",
			HTTPCode:  http.StatusOK,
			Datetime:  time.Now().Format("2006-01-02 15:04:05"),
			Timestamp: time.Now().Unix(),
			User:      users,
		}, http.StatusOK)
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type (
		request struct {
			Name      string `json:"name"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Email     string `json:"email"`
			Phone     string `json:"phone"`
			Age       int32  `json:"age"`
		}

		response struct {
			Status    string      `json:"status"`
			HTTPCode  int         `json:"http_code"`
			Datetime  string      `json:"datetime"`
			Timestamp int64       `json:"timestamp"`
			User      interface{} `json:"user"`
		}
	)
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req request
			err error
		)

		if err = s.decode(r, &req); err != nil {
			s.config.logger.Errorf("invalid user request: %v", err)
			s.respondErr(w, BadRequest())
			return
		}

		params := &CreateUserParams{
			Name:      req.Name,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
			Age:       req.Age,
		}

		createdUser, err := s.config.userStore.CreateUser(r.Context(), params)
		if err != nil {
			s.config.logger.Errorf("error creating user:: %v", err)
			s.respondErr(w, InternalServerError())
			return
		}

		s.respond(w, response{
			Status:    "success",
			HTTPCode:  http.StatusCreated,
			Datetime:  time.Now().Format("2006-01-02 15:04:05"),
			Timestamp: time.Now().Unix(),
			User:      createdUser,
		}, http.StatusCreated)
	}
}
