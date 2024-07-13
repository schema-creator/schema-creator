package dai

import (
	"context"
)

type DataAccessInterfaces interface {
	UserRepo
	SessionRepo

	Transaction(ctx context.Context, fn func(context.Context, DataAccessInterfaces) error) error
}
