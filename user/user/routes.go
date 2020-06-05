package user

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/orderforme/user/middelware"
)

// Register a without JWT
func (handler *Handler) Register(route *echo.Group) {

	user := route.Group("/user")
	user.POST("/login", handler.LoginUser)
	user.POST("/create", handler.CreateUser)
	//only use with algorithm RS256
	user.POST("/validate", handler.ValidateToken)
}

// Register a with JWT
func (handler *Handler) RegisterJWT(route *echo.Group) {

	user := route.Group("/user")
	user.GET("", handler.GetUserByID)
	user.GET("/total", handler.GetUsers)
	user.PUT("", handler.UpdateUser)
	user.PUT("/avatar", handler.UploadAvatar)
	user.PUT("/front", handler.UploadFront)
	user.DELETE("", handler.DeleteUser)
	user.POST("/hola", hola)
}

func hola(context echo.Context) error {
	hola := context.Get("userid").(*jwt.Token)
	claims := hola.Claims.(*middelware.Claim)
	hola11 := claims.UserName
	return context.JSON(http.StatusOK, hola11)

}
