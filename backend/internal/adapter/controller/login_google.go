package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/cookie"
	"github.com/schema-creator/schema-creator/schema-creator/internal/framework/hcontext"
	"github.com/schema-creator/schema-creator/schema-creator/internal/usecase/interactor"
	"github.com/schema-creator/schema-creator/schema-creator/pkg/log"
)

type LoginRequest struct {
	AuthorizationCode string `json:"code"`
}

func (r LoginRequest) Validate() error {
	if r.AuthorizationCode == "" {
		return fmt.Errorf("authorization code is required")
	}
	return nil
}

type LoginResponse struct {
	UserID string `json:"userId"`
}

func GoogleLogin(
	googleLogin *interactor.GoogleLogin,
	cookieSetter *cookie.CookieSetter,
) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		var reqBody LoginRequest
		if err := ctx.Bind(&reqBody); err != nil {
			log.Error(ctx.Request().Context(), "bad request", err)
			return echo.ErrBadRequest
		}
		fmt.Println(reqBody)
		if err := reqBody.Validate(); err != nil {
			log.Error(ctx.Request().Context(), "bad request", err)
			return echo.ErrBadRequest
		}

		userAgent, ok := ctx.Get(hcontext.UserAgent.String()).(string)
		if !ok {
			log.Error(ctx.Request().Context(), "user agent not found")
			return echo.ErrInternalServerError
		}

		userInfo, err := googleLogin.GoogleLogin(ctx.Request().Context(), reqBody.AuthorizationCode, userAgent)
		if err != nil {
			log.Error(ctx.Request().Context(), "failed to google login", err)
			return echo.ErrInternalServerError
		}

		cookieSetter.CreateCookieSetter(ctx).SetCookieValue(string(cookie.SessionID), userInfo.SessionID)

		return ctx.JSON(http.StatusOK, &LoginResponse{
			UserID: userInfo.UserID,
		})
	}
}
