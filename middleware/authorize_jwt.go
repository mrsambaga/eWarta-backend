package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AppError struct {
	StatusCode int    `json:"statusCode"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func UnauthorizedError() AppError {
	return AppError{
		Code:       "UNAUTHORIZED_ERROR",
		Message:    "Unauthorize",
		StatusCode: http.StatusUnauthorized,
	}
}

func (err AppError) Error() string {
	return err.Message
}

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}

func AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	s := strings.Split(authHeader, " ")
	authError := UnauthorizedError()
	if len(s) < 2 {
		c.AbortWithStatusJSON(authError.StatusCode, authError)
		return
	}

	tokenString := s[1]
	token, err := validateToken(tokenString)
	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(authError.StatusCode, authError)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(authError.StatusCode, authError)
		return
	}

	c.Set("id", int(claims["id"].(float64)))
	c.Next()
}
