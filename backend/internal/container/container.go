package container

import (
	"github.com/schema-creator/schema-creator/schema-creator/internal/adapter/gateway/authz"
	"github.com/schema-creator/schema-creator/schema-creator/internal/adapter/gateway/authz/github"
	"github.com/schema-creator/schema-creator/schema-creator/internal/adapter/gateway/authz/google"
	"github.com/schema-creator/schema-creator/schema-creator/internal/adapter/gateway/repository"
	"github.com/schema-creator/schema-creator/schema-creator/internal/driver/db"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/cookie"
	"github.com/schema-creator/schema-creator/schema-creator/internal/usecase/dai"
	"github.com/schema-creator/schema-creator/schema-creator/internal/usecase/interactor"
	"go.uber.org/dig"
)

var container *dig.Container

type provideArg struct {
	constructor any
	opts        []dig.ProvideOption
}

func NewContainer() error {
	container = dig.New()

	args := []provideArg{
		{constructor: google.DefaultGoogleOAuth2Config, opts: []dig.ProvideOption{}},
		{constructor: github.DefaultGitHubOAuth2Config, opts: []dig.ProvideOption{}},
		{constructor: db.Connect, opts: []dig.ProvideOption{}},
		{constructor: db.NewGORM, opts: []dig.ProvideOption{}},
		{constructor: cookie.DefaultCookieOptions, opts: []dig.ProvideOption{}},
		{constructor: cookie.NewCookieSetter, opts: []dig.ProvideOption{}},
		{constructor: google.NewGoogleOAuth, opts: []dig.ProvideOption{dig.As(new(authz.GoogleOAuth2))}},
		{constructor: github.NewGitHubOAuth, opts: []dig.ProvideOption{dig.As(new(authz.GitHubOAuth2))}},
		{constructor: interactor.NewGoogleLogin, opts: []dig.ProvideOption{}},
		{constructor: repository.NewGormRepo, opts: []dig.ProvideOption{dig.As(new(dai.DataAccessInterfaces))}},
		{constructor: interactor.NewGitHubLogin, opts: []dig.ProvideOption{}},
		{constructor: interactor.NewSessionInteractor, opts: []dig.ProvideOption{}},
	}

	for _, arg := range args {
		if err := container.Provide(arg.constructor, arg.opts...); err != nil {
			return err
		}
	}

	return nil
}

func Invoke[T any]() T {
	var r T
	if err := container.Invoke(func(t T) error {
		r = t
		return nil
	}); err != nil {
		panic(err)
	}

	return r
}
