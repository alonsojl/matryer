package apirest

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type dbUser struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
	Age       int32  `db:"age"`
}

type UserStore interface {
	GetUsers() ([]*UserData, error)
	CreateUser(context.Context, *CreateUserParams) (*UserData, error)
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

func (s *MySQLUserStore) GetUsers() ([]*UserData, error) {
	var dbUsers []*dbUser
	err := s.client.Select(&dbUsers, "CALL spGetUsers()")
	if err != nil {
		s.logger.Errorf("error get users: %v", err)
		return nil, err
	}

	if len(dbUsers) == 0 {
		s.logger.Error("there are no user records")
		return nil, err
	}

	var users []*UserData
	for _, dbUser := range dbUsers {
		users = append(users, &UserData{
			ID:        dbUser.ID,
			Name:      dbUser.Name,
			FirstName: dbUser.FirstName,
			LastName:  dbUser.LastName,
			Email:     dbUser.Email,
			Phone:     dbUser.Phone,
			Age:       dbUser.Age,
		})
	}

	return users, nil
}

func (s *MySQLUserStore) CreateUser(ctx context.Context, user *CreateUserParams) (*UserData, error) {
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

	return &UserData{
		ID:        userID,
		Name:      user.Name,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Age:       user.Age,
	}, nil
}
