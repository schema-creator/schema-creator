package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/cookie"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/hcontext"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/herror"
	"github.com/schema-creator/schema-creator/schema-creator/internal/usecase/interactor"
	"github.com/schema-creator/schema-creator/schema-creator/pkg/log"
)

func SessionMiddleware(ui *interactor.SessionInteractor) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			sessionID, err := ctx.Cookie(string(cookie.SessionID))
			if err != nil {
				if err == http.ErrNoCookie {
					log.Info(ctx.Request().Context(), "session cookie not found")
					return echo.ErrUnauthorized
				}
				log.Error(ctx.Request().Context(), fmt.Sprintf("failed to get session cookie: %v", err))
				return echo.ErrInternalServerError
			}

			if sessionID.Value == "" {
				log.Info(ctx.Request().Context(), "session cookie value is empty")
				return echo.ErrUnauthorized
			}

			userAgent, ok := ctx.Get(hcontext.UserAgent.String()).(string)
			if !ok {
				log.Info(ctx.Request().Context(), "user agent not found")
				return echo.ErrBadRequest
			}
			user, err := ui.Repositories.GetSessionByID(ctx.Request().Context(), sessionID.Value, userAgent)

			if err != nil {
				log.Info(ctx.Request().Context(), fmt.Sprintf("failed to get user by session id: %v", err))
				switch err {
				case herror.ErrSessionExpired:
					return echo.ErrUnauthorized
				case herror.ErrResourceNotFound:
					return echo.ErrUnauthorized
				default:
					return echo.ErrInternalServerError
				}
			}
			ctx.Set(hcontext.UserID.String(), user.UserID)
			return next(ctx)
		}
	}
}
