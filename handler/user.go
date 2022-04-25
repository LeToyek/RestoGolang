package handler

import (
	"net/http"
	"resto/entities"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertUser(c *gin.Context) {
	var userReq entities.User
	if err := c.BindJSON(&userReq); err != nil {
		panic(err)
	}
	_, err := h.Service.AddUser(userReq)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"created": userReq})
}
