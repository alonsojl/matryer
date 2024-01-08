package api

import (
	"matryer/model"
	"net/http"
	"time"
)

type userResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone,omitempty"`
	Age       int32  `json:"age"`
}

func (s *server) handleUsersGet() http.HandlerFunc {
	type response struct {
		Status    string         `json:"status"`
		HTTPCode  int            `json:"http_code"`
		Datetime  string         `json:"datetime"`
		Timestamp int64          `json:"timestamp"`
		User      []userResponse `json:"user"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := s.config.userStore.GetUsers()
		if err != nil {
			s.config.logger.Errorf("error getting users: %v", err)
			s.renderResponse(w, err)
			return
		}

		var userResp []userResponse
		for _, user := range users {
			userResp = append(userResp, userResponse{
				ID:        user.ID,
				Name:      user.Name,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Phone:     user.Phone,
				Age:       user.Age,
			})
		}

		s.respond(w, response{
			Status:    "success",
			HTTPCode:  http.StatusOK,
			Datetime:  time.Now().Format("2006-01-02 15:04:05"),
			Timestamp: time.Now().Unix(),
			User:      userResp,
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
			s.renderResponse(w, BadRequest())
			return
		}

		user, err := s.config.userStore.CreateUser(r.Context(), &model.UserParams{
			Name:      req.Name,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
			Age:       req.Age,
		})

		if err != nil {
			s.config.logger.Errorf("error creating user:: %v", err)
			s.renderResponse(w, InternalServerError())
			return
		}

		s.respond(w, response{
			Status:    "success",
			HTTPCode:  http.StatusCreated,
			Datetime:  time.Now().Format("2006-01-02 15:04:05"),
			Timestamp: time.Now().Unix(),
			User: userResponse{
				ID:        user.ID,
				Name:      user.Name,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Phone:     user.Phone,
				Age:       user.Age,
			},
		}, http.StatusCreated)
	}
}
