package dai

import (
	"context"
)

type DataAccessInterfaces interface {
	Transaction(ctx context.Context, fn func(context.Context, DataAccessInterfaces) error) error
}
