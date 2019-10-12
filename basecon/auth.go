package basecon

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tString := c.Request.Header.Get("Authorization")
	t, err := jwt.Parse(tString, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != t.Method {
			return nil, fmt.Errorf("Wrong signing method: %v", t.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if t != nil && err == nil {
		fmt.Println("token verified")
	} else {
		r := gin.H{
			"message": "Not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, r)
		c.Abort()
	}
}
