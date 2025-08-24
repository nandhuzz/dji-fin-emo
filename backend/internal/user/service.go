package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nandhuzz/dji-fin-emo/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(ctx context.Context, name, email, password string) (*User, error)
	Login(ctx context.Context, email, password string) (string, error)
	Details(ctx context.Context, userId string) (*User, error)
}

type service struct {
	repo UserRepository
}

func NewService(repo UserRepository) Service {
	return &service{repo: repo}
}

func (s *service) Register(ctx context.Context, name, email, password string) (*User, error) {
	existing, _ := s.repo.GetUserByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("email already in use")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user := User{
		ID:        uuid.NewString(),
		Name:      name,
		Email:     email,
		Password:  string(hashed),
		CreatedAt: time.Now(),
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil || user == nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	return auth.GenerateToken(user.ID)
}

func (s *service) Details(ctx context.Context, userId string) (*User, error) {
	user, err := s.repo.GetUserByID(ctx, userId)
	if err != nil || user == nil {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}
