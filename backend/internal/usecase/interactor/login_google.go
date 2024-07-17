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

type GoogleLogin struct {
	authz        authz.GoogleOAuth2
	repositories dai.DataAccessInterfaces
}

func NewGoogleLogin(
	authz authz.GoogleOAuth2,
	repositories dai.DataAccessInterfaces,
) *GoogleLogin {
	return &GoogleLogin{
		authz:        authz,
		repositories: repositories,
	}
}

type LoginResult struct {
	SessionID string
	UserID    string
	Icon      string
}

func (gl *GoogleLogin) GoogleLogin(ctx context.Context, authorizationCode, userAgent string) (*LoginResult, error) {
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

	if err := gl.repositories.Transaction(ctx, func(ctx context.Context, da dai.DataAccessInterfaces) error {
		// 居ないなら作る
		if !found {
			_, err := gl.repositories.CreateUser(ctx, &model.User{
				UserID:    userInfo.UserID,
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

		// Tokenを更新
		_, err = gl.repositories.SyncSession(ctx, &model.Session{
			SessionID:      sessionID.String(),
			UserAgent:      userAgent,
			UserID:         userInfo.UserID,
			Token:          token.AccessToken,
			ExpirationTime: int32(token.Expiry.Unix()),
		})

		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &LoginResult{
		SessionID: sessionID.String(),
		UserID:    userInfo.UserID,
		Icon:      userInfo.Icon,
	}, nil
}
