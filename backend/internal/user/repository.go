package user

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type UserRepository interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByID(ctx context.Context, userId string) (*User, error)
}

type repository struct {
	tx pgx.Tx
}

func NewRepository(tx pgx.Tx) UserRepository {
	return &repository{tx: tx}
}

func (r *repository) CreateUser(ctx context.Context, u User) error {
	_, err := r.tx.Exec(ctx, `
		INSERT INTO users (id, name, email, password_hash, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`, u.ID, u.Name, u.Email, u.Password, u.CreatedAt)
	return err
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	row := r.tx.QueryRow(ctx, `
		SELECT id, name, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`, email)

	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *repository) GetUserByID(ctx context.Context, userId string) (*User, error) {
	row := r.tx.QueryRow(ctx, `
		SELECT id, name, email, password_hash, created_at
		FROM users
		WHERE id = $1
	`, userId)

	var u User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
