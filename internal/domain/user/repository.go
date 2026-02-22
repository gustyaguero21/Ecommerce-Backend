package user

import (
	"context"
)

type Repository interface {
	GetUserByID(ctx context.Context, userID string) (*User, error)
}
