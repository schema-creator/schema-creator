package github

import (
	"context"

	"github.com/schema-creator/schema-creator/schema-creator/cmd/config"
	"github.com/schema-creator/schema-creator/schema-creator/internal/adapter/gateway/authz"
	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type OAuth2 struct {
	oac oauth2.Config
}

func DefaultOAuth2Config() oauth2.Config {
	return oauth2.Config{
		ClientID:     config.Config.Github.ClientID,
		ClientSecret: "github_client_secret",
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
		RedirectURL: "",
	}
}

func NewOAuth(oac oauth2.Config) *OAuth2 {
	return &OAuth2{
		oac: oac,
	}
}

func (o *OAuth2) FetchToken(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := o.oac.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (o *OAuth2) GetUserInfo(ctx context.Context, token *oauth2.Token) (*authz.UserInfo, error) {
	client := o.oac.Client(ctx, token)

	service, err := v2.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	gitHubUser, err := service.Userinfo.V2.Me.Get().Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return &authz.UserInfo{
		UserID: gitHubUser.Id,
		Email:  gitHubUser.Email,
		Icon:   gitHubUser.Picture,
	}, nil
}

var _ authz.OAuth2 = (*OAuth2)(nil)
