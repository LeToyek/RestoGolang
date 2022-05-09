package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Validation() {

}

func midd() gin.HandlerFunc {
	return func(c *gin.Context) {

		auth := c.Request.Header["Authorization"]
		fmt.Println(auth)
		if len(auth) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
				"success": false,
				"message": "Auth required",
			})
			return
		}
		token := strings.Split(auth[0], " ")[1]
		if token != "toyek.goblog" {
			c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
				"success": false,
				"message": "Auth required",
			})
			return
		}
		c.Set("UserToken", token)
		c.Next()
	}
}
