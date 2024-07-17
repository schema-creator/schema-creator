package google

import (
	"context"

	"github.com/schema-creator/schema-creator/schema-creator/cmd/config"
	"github.com/schema-creator/schema-creator/schema-creator/internal/adapter/gateway/authz"
	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"

	"google.golang.org/api/option"
)

type GoogleOAuth2 struct {
	oac oauth2.Config
}
type GoogleOAuth2Config struct {
	*oauth2.Config
}

func DefaultGoogleOAuth2Config() GoogleOAuth2Config {
	return GoogleOAuth2Config{
		Config: &oauth2.Config{
			ClientID:     config.Config.Google.ClientID,
			ClientSecret: config.Config.Google.ClientSecret,
			Scopes:       []string{},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.google.com/o/oauth2/auth",
				TokenURL: "https://accounts.google.com/o/oauth2/token",
			},
			RedirectURL: config.Config.Google.RedirectURI,
		},
	}

}

func NewGoogleOAuth(oac GoogleOAuth2Config) *GoogleOAuth2 {
	return &GoogleOAuth2{
		oac: *oac.Config,
	}
}

func (o *GoogleOAuth2) FetchToken(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := o.oac.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (o *GoogleOAuth2) GetUserInfo(ctx context.Context, token *oauth2.Token) (*authz.UserInfo, error) {
	client := o.oac.Client(ctx, token)

	service, err := v2.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	googleUser, err := service.Userinfo.V2.Me.Get().Context(ctx).Do()
	if err != nil {
		return nil, err
	}

	return &authz.UserInfo{
		UserID: googleUser.Id,
		Email:  googleUser.Email,
		Icon:   googleUser.Picture,
	}, nil
}

var _ authz.GoogleOAuth2 = (*GoogleOAuth2)(nil)
