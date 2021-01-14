package middlewares

import (
	"net/http"
	"strings"

	infraAuth "github.com/martinsd3v/simple-golang-project/infra/auth"

	"github.com/gin-gonic/gin"
)

//AuthMiddleware responsible for check JTW tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := extractToken(c.Request)
		err := infraAuth.TokenValid(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

//get the token from the request body
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	simpleToken := r.Header.Get("token")
	if simpleToken != "" {
		return simpleToken
	}

	return ""
}
