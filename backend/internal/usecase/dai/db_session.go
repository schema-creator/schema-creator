package dai

import (
	"context"

	"github.com/schema-creator/schema-creator/schema-creator/internal/entities/model"
)

type SessionRepo interface {
	GetSessionByID(ctx context.Context, sessionID, userAgent string) (*model.Session, error)
	SyncSession(ctx context.Context, session *model.Session) (*model.Session, error)
}
