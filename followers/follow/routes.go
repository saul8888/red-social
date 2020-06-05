package follow

import (
	"github.com/labstack/echo"
)

// Register a new user
func (handler *Handler) RegisterJWT(route *echo.Group) {

	follow := route.Group("/follow")
	//follow.GET("", handler.Followers)
	follow.POST("", handler.Follow)
	follow.DELETE("", handler.Unfollow)
	follow.DELETE("", handler.Queryfollow)

}
