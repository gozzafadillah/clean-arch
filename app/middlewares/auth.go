package middlewares

import (
	"errors"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID     int  `json:"id"`
	Status bool `json:"status"`
	jwt.StandardClaims
}

type ConfigJwt struct {
	SecretJWT string
}

func (jwtConf *ConfigJwt) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return errors.New("cannot generete a token")
		}),
	}
}

// GenerateToken jwt ...
func (jwtConf *ConfigJwt) GenerateToken(userID int, userStatus bool) (string, error) {
	claims := JwtCustomClaims{
		ID:     userID,
		Status: userStatus,
	}

	// Create token with claims
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(jwtConf.SecretJWT))

	return token, err
}

// GetUser from jwt ...
func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}
