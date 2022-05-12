package server

import (
	"resto/handler"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router  *gin.Engine
	Handler *handler.Handler
}

func (s *Server) StartServer() {
	s.Router.POST("/register", s.Handler.InsertUser)
	s.Router.GET("/user/:id", s.Handler.GetUser)
	s.Router.GET("/users", s.Handler.GetUsers)
	s.Router.POST("/login", s.Handler.Login)
	s.Router.GET("/logout", s.Handler.Logout)
	s.Router.POST("/refresh", s.Handler.RefreshAccountToken)
	s.Router.POST("/table", s.Handler.AddTable)
	s.Router.GET("/tables", s.Handler.GetAllTables)
	s.Router.POST("/food", s.Handler.AddFood)
	s.Router.GET("/foods", s.Handler.GetAllfoods)
	s.Router.GET("/foods/:category", s.Handler.GetFoodsByCategory)
	s.Router.POST("/order", s.Handler.AddOrder)
	s.Router.GET("/order", s.Handler.GetOrder)
	// s.Router.GET("/test", midd(), func(ctx *gin.Context) {
	// 	userToken := ctx.GetString("UserToken")
	// 	ctx.IndentedJSON(http.StatusOK, map[string]interface{}{
	// 		"success":    true,
	// 		"message":    "Hello World",
	// 		"user_token": userToken,
	// 	})
	// })
}

// Auth middleware
