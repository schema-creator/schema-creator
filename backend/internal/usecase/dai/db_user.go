package dai

import (
	"context"

	"github.com/schema-creator/schema-creator/schema-creator/internal/entities/model"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByID(ctx context.Context, userID string) (*model.User, error)
	GetUsers(ctx context.Context, limit, offset int) ([]model.User, error)
}
