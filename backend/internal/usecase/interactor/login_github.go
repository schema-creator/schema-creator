package interactor

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/schema-creator/schema-creator/schema-creator/internal/adapter/gateway/authz"
	"github.com/schema-creator/schema-creator/schema-creator/internal/entities/model"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/herror"
	"github.com/schema-creator/schema-creator/schema-creator/internal/usecase/dai"
)

type GitHubLogin struct {
	authz        authz.OAuth2
	repositories dai.DataAccessInterfaces
}

func NewGitHubLogin(
	authz authz.OAuth2,
	repositories dai.DataAccessInterfaces,
) *GitHubLogin {
	return &GitHubLogin{
		authz:        authz,
		repositories: repositories,
	}
}

// type LoginResult struct {
// 	SessionID string
// 	UserID    string
// 	Icon      string
// }

func (gl *GitHubLogin) GitHubLogin(ctx context.Context, authorizationCode, userAgent string) (*LoginResult, error) {
	token, err := gl.authz.FetchToken(ctx, authorizationCode)
	if err != nil {
		return nil, err
	}

	userInfo, err := gl.authz.GetUserInfo(ctx, token)
	if err != nil {
		return nil, err
	}

	var found bool = true
	// userが存在するかチェック
	_, err = gl.repositories.GetUserByID(ctx, userInfo.UserID)
	if err != nil {
		if !errors.Is(err, herror.ErrResourceNotFound) {
			return nil, nil
		}
		found = false
	}

	sessionID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	if err := gl.repositories.Transaction(ctx, func(ctx context.Context, tx dai.DataAccessInterfaces) error {
		if !found {
			user := &model.User{
				UserID: userInfo.UserID,
				Icon:   userInfo.Icon,
			}
			if _, err := tx.UserRepo.SyncUser(ctx, user); err != nil {
				return err
			}
		}

		session := &model.Session{
			SessionID:      sessionID.String(),
			UserID:         userInfo.UserID,
			UserAgent:      userAgent,
			ExpirationTime: int32(token.Expiry.Unix()),
		}
		if _, err := tx.SessionRepo.SyncSession(ctx, session); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
}
