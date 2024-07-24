package authz

import (
	"context"

	"golang.org/x/oauth2"
)

type GoogleOAuth2 interface {
	FetchToken(ctx context.Context, authorizationCode string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, token *oauth2.Token) (*UserInfo, error)
}
type GitHubOAuth2 interface {
	FetchToken(ctx context.Context, authorizationCode string) (*oauth2.Token, error)
	GetUserInfo(ctx context.Context, token *oauth2.Token) (*UserInfo, error)
}

type UserInfo struct {
	UserID string
	Name   string
	Email  string
	Icon   string
}
