package models

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context) ([]*models.User, error)
}
