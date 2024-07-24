package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/schema-creator/schema-creator/schema-creator/cmd/config"
	"github.com/schema-creator/schema-creator/schema-creator/internal/adapter/gateway/authz"
	"golang.org/x/oauth2"
)

type GitHubOAuth2 struct {
	oac oauth2.Config
}
type GitHubOAuth2Config struct {
	*oauth2.Config
}
type GitHubUser struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	LoginName string `json:"login"`
	Email     string `json:"email"`
	Icon      string `json:"avatar_url"`
}
type GitHubEmail struct {
	Email   string `json:"email"`
	Primary bool   `json:"primary"`
}

func DefaultGitHubOAuth2Config() GitHubOAuth2Config {
	return GitHubOAuth2Config{
		Config: &oauth2.Config{
			ClientID:     config.Config.Github.ClientID,
			ClientSecret: config.Config.Github.ClientSecret,
			Scopes:       []string{"read:user", "user:email"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://github.com/login/oauth/authorize",
				TokenURL: "https://github.com/login/oauth/access_token",
			},
			RedirectURL: config.Config.Github.RedirectURI,
		},
	}
}

func NewGitHubOAuth(oac GitHubOAuth2Config) *GitHubOAuth2 {
	return &GitHubOAuth2{
		oac: *oac.Config,
	}
}

func (o *GitHubOAuth2) FetchToken(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := o.oac.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (o *GitHubOAuth2) GetUserInfo(ctx context.Context, token *oauth2.Token) (*authz.UserInfo, error) {
	client := o.oac.Client(ctx, token)
	res, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	userInfo, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, echo.ErrBadRequest
	}

	var gitHubUser GitHubUser
	var gitHubEmails []GitHubEmail
	if err := json.Unmarshal(userInfo, &gitHubUser); err != nil {
		return nil, err
	}
	if gitHubUser.Name == "" {
		gitHubUser.Name = gitHubUser.LoginName
	}

	if gitHubUser.Email == "" {
		emailRes, err := client.Get("https://api.github.com/user/emails")
		if err != nil {
			return nil, err
		}
		defer emailRes.Body.Close()
		emailInfo, err := io.ReadAll(emailRes.Body)
		if err != nil {
			return nil, err
		} else if emailRes.StatusCode != 200 {
			return nil, echo.ErrBadRequest
		}
		if err := json.Unmarshal(emailInfo, &gitHubEmails); err != nil {
			return nil, err
		}
		for _, email := range gitHubEmails {
			if email.Primary {
				gitHubUser.Email = email.Email
				break
			}
		}
	}

	return &authz.UserInfo{
		UserID: fmt.Sprintf("%d", gitHubUser.ID),
		Name:   gitHubUser.Name,
		Email:  gitHubUser.Email,
		Icon:   gitHubUser.Icon,
	}, nil
}

var _ authz.GitHubOAuth2 = (*GitHubOAuth2)(nil)
