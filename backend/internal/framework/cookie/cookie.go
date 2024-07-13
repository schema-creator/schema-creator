package cookie

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CookieOptions struct {
	MaxAge   int
	Path     string
	SameSite http.SameSite
	Secure   bool
	Httponly bool
}

func DefaultCookieOptions() CookieOptions {
	return CookieOptions{
		MaxAge:   int(time.Now().Unix()) + 10*365*24*3600,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
		Httponly: true,
	}
}

type CookieSetter struct {
	CookieOptions CookieOptions
}

func NewCookieSetter(option CookieOptions) *CookieSetter {
	return &CookieSetter{
		CookieOptions: option,
	}
}
func (f *CookieSetter) CreateCookieSetter(c echo.Context) *EchoCookieSetter {
	setter := NewEchoCookieSetter(c, f.CookieOptions)
	return setter
}

type EchoCookieSetter struct {
	origin         string
	ctx            echo.Context
	defaultOptions CookieOptions
}

func NewEchoCookieSetter(c echo.Context, defaultOptions CookieOptions) *EchoCookieSetter {
	return &EchoCookieSetter{
		origin:         c.Request().Header.Get("Origin"),
		ctx:            c,
		defaultOptions: defaultOptions,
	}
}

func (ecs *EchoCookieSetter) SetCookie(name, value string, maxAge int, path string, secure, httponly bool) {
	domain := ""
	ecs.ctx.SetCookie(&http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: ecs.defaultOptions.SameSite,
		Secure:   secure,
		HttpOnly: httponly,
	})
}

func (ecs *EchoCookieSetter) SetCookieValue(name, value string) {
	ecs.SetCookie(name, value, ecs.defaultOptions.MaxAge, ecs.defaultOptions.Path, ecs.defaultOptions.Secure, ecs.defaultOptions.Httponly)
}
