package models

import "context"

type Service interface {
	CreateUser(ctx context.Context, requesterID, login, password, role string) (*models.User, error)
	Authenticate(ctx context.Context, login, password string) (*models.User, error)
	DeleteUser(ctx context.Context, requesterID, userID string) error
	ListUsers(ctx context.Context, requesterID string) ([]models.User, error)
}
