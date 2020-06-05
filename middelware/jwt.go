package middelware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

//authentication
var autoConfig = middleware.JWTConfig{
	Claims:        &Claim{},
	SigningMethod: jwt.SigningMethodHS256.Name,
	SigningKey:    Keys(),
	ContextKey:    "userid",
	//SigningMethod: jwt.SigningMethodRS256.Name,
	//SigningKey:    authentication.PublicKey,
	//TokenLookup: "header:" + echo.HeaderAuthorization,
}

//Bearer {token}
