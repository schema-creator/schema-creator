package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

type router struct {
	echo *echo.Echo
}

func NewRouter() http.Handler {
	echo := echo.New()
	router := &router{
		echo: echo,
	}

	// setup middleware
	router.echo.Use(echoMiddleware.Recover())

	// router.echo.Use(echoprometheus.NewMiddleware("hal-cinema"))

	router.health()

	corsRoute := router.echo.Group("")

	corsRoute.Use(echoMiddleware.CORSWithConfig(echoMiddleware.DefaultCORSConfig))

	return router.echo
}

func (r *router) health() {
	r.echo.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"status:":"ok"}`)
	})
}
