package server

import (
	"fmt"
	"net/http"
	"resto/handler"
	"strings"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Handler *handler.Handler
}

func (s *Server) StartServer() {
	s.Router.POST("/register", s.Handler.InsertUser)
	s.Router.GET("/test", midd(), func(ctx *gin.Context) {
		userToken := ctx.GetString("UserToken")
		ctx.IndentedJSON(http.StatusOK, map[string]interface{}{
			"success":    true,
			"message":    "Hello World",
			"user_token": userToken,
		})
	})
}

// Auth middleware
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
