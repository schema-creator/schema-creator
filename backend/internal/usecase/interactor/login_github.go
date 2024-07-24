package interactor

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/schema-creator/schema-creator/schema-creator/internal/adapter/gateway/authz"
	"github.com/schema-creator/schema-creator/schema-creator/internal/entities/model"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/herror"
	"github.com/schema-creator/schema-creator/schema-creator/internal/usecase/dai"
)

type GitHubLogin struct {
	authz        authz.GitHubOAuth2
	repositories dai.DataAccessInterfaces
}

func NewGitHubLogin(
	authz authz.GitHubOAuth2,
	repositories dai.DataAccessInterfaces,
) *GitHubLogin {
	return &GitHubLogin{
		authz:        authz,
		repositories: repositories,
	}
}

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
			_, err := gl.repositories.CreateUser(ctx, &model.User{
				UserID:    userInfo.UserID,
				Name:      userInfo.Name,
				Email:     userInfo.Email,
				Icon:      userInfo.Icon,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				IsDeleted: false,
			})

			if err != nil {
				return err
			}
		}

		session := &model.Session{
			SessionID:      sessionID.String(),
			UserID:         userInfo.UserID,
			UserAgent:      userAgent,
			Token:          token.AccessToken,
			ExpirationTime: int32(time.Now().Add(time.Hour * 24 * 30).Unix()),
		}
		if _, err := gl.repositories.SyncSession(ctx, session); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return &LoginResult{
		SessionID: sessionID.String(),
		UserID:    userInfo.UserID,
		Name:      userInfo.Name,
		Email:     userInfo.Email,
		Icon:      userInfo.Icon,
	}, nil
}
