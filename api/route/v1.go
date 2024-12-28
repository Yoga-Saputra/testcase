package route

import (
	"github.com/labstack/echo/v4"
)

func v1(e *echo.Echo) *echo.Group {
	return e.Group("/v1")
}
