package public

import (
	"github.com/labstack/echo"
)

// Register a new user
func (handler *Handler) RegisterJWT(route *echo.Group) {

	public := route.Group("/public")
	public.GET("", handler.GetPublicByID)
	public.GET("/total", handler.GetPublics)
	public.POST("", handler.CreatePublic)
	public.PUT("", handler.UpdatePublic)
	public.DELETE("", handler.DeletePublic)
	public.POST("/token", handler.ValidateToken)

}
