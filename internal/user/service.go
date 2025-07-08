package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Service interface {
	Register(ctx context.Context, name, email, password string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(ctx context.Context, name, email, password string) (*User, error) {
	existing, _ := s.repo.GetUserByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("email already in use")
	}

	user := User{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  password, // TODO: hash password
		CreatedAt: time.Now(),
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
