package api

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserStore interface {
	GetUsers() ([]*User, error)
	CreateUser(context.Context, *CreateUserParams) (*User, error)
}

type MySQLUserStore struct {
	logger *logrus.Logger
	client *sqlx.DB
}

func NewMySQLUserStore(client *sqlx.DB, logger *logrus.Logger) *MySQLUserStore {
	return &MySQLUserStore{
		client: client,
		logger: logger,
	}
}

func (s *MySQLUserStore) GetUsers() ([]*User, error) {
	var users []*User
	err := s.client.Select(&users, "CALL spGetUsers()")
	if err != nil {
		s.logger.Errorf("error get users: %v", err)
		return nil, err
	}

	if len(users) == 0 {
		s.logger.Error("there are no user records")
		return nil, err
	}

	return users, nil
}

func (s *MySQLUserStore) CreateUser(ctx context.Context, user *CreateUserParams) (*User, error) {
	var userID int64
	query := "CALL spCreateUser(?,?,?,?,?,?)"
	err := s.client.QueryRowxContext(ctx, query,
		user.Name,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Phone,
		user.Age,
	).Scan(&userID)

	if err != nil {
		s.logger.Errorf("failed to insert user: %v", err)
		return nil, err
	}

	return &User{
		ID:        userID,
		Name:      user.Name,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Age:       user.Age,
	}, nil
}
