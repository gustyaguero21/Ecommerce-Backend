package user

import (
	"context"
	"ecommerce-backend/internal/domain/user"
	errors "ecommerce-backend/internal/domain/user"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	Pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) user.Repository {
	return &UserRepository{Pool: pool}
}

func (ur *UserRepository) GetUserByID(ctx context.Context, userID string) (*user.User, error) {

	user := &user.User{}

	err := ur.Pool.QueryRow(ctx, GetUserByIDQuery, userID).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.DNI,
		&user.Email,
		&user.Telephone,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
