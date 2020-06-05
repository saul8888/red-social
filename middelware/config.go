package middelware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Middelware(route *echo.Group) {
	route.Use(middleware.LoggerWithConfig(logConfig))
	route.Use(middleware.RecoverWithConfig(recoConfig))
	route.Use(middleware.CORSWithConfig(corsConfig))
}

func MiddelwareJWT(route *echo.Group) {
	route.Use(middleware.JWTWithConfig(autoConfig))
}
