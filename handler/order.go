package handler

import (
	"net/http"
	"resto/entities"
	"time"

	helper "resto/helper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddOrder(c *gin.Context) {
	claims, err := helper.ValidateToken(c)

	if err != nil {
		panic(err)
	}
	var order entities.Order

	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, err.Error())
	}

	order.User_id = claims.ID
	order.Created_at = time.Now().String()
	order.Updated_at = time.Now().String()

	err = h.Service.AddOrder(order)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		panic(err)

	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"order": order,
	})
}

func (h *Handler) GetOrder(c *gin.Context) {
	claims, err := helper.ValidateToken(c)

	if err != nil {
		panic(err)
	}

	userID := claims.ID
	order, err := h.Service.GetOrder(userID)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"order": order,
	})

}
