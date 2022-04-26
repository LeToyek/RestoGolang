package handler

import (
	"net/http"
	"resto/entities"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/gin-gonic/gin"
)

func (h *Handler) InsertUser(c *gin.Context) {
	var userReq entities.User
	if err := c.BindJSON(&userReq); err != nil {
		panic(err)
	}
	// countEmail := fmt.Sprintf()
	userReq.User_id = nanoid.New()
	userReq.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	userReq.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	_, err := h.Service.AddUser(userReq)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"created": userReq})
}
