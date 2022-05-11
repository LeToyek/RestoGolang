package handler

import (
	"fmt"
	"net/http"
	"resto/entities"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddFood(c *gin.Context) {
	var food entities.Food

	if err := c.BindJSON(&food); err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"errors": err.Error(),
		})
	}

	food.Food_id = nanoid.New()
	food.Created_at = time.Now().String()
	food.Updated_at = time.Now().String()
	err := h.Service.AddFood(food)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"errors": err.Error(),
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"created": food,
	})
}

func (h *Handler) GetAllfoods(c *gin.Context) {
	foods, err := h.Service.GetAllfoods()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"errors": err.Error(),
		})
	}

	c.IndentedJSON(http.StatusOK, foods)
}

func (h *Handler) GetFoodsByCategory(c *gin.Context) {
	category := c.Param("category")
	fmt.Println("---->" + category)

	foods, err := h.Service.GetFoodsByCategory(category)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, gin.H{
			"errors": err.Error(),
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		category: foods,
	})
}
