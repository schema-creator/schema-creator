package v1

import (
	"github.com/labstack/echo/v4"
)

type v1Router struct {
	engine *echo.Group
}

func Setup(engine *echo.Group) {
	v1 := &v1Router{
		engine: engine,
	}

	{
		v1.engine.POST("/login/google", nil)
	}
}
