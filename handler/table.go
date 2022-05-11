package handler

import (
	"net/http"
	"resto/entities"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddTable(c *gin.Context) {
	var table entities.TableResto

	if err := c.BindJSON(&table); err != nil {
		panic(err)
	}

	table.Created_at = time.Now().String()
	table.Updated_at = time.Now().String()

	err := h.Service.AddTable(table)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"status":  "success",
		"created": table,
	})
}
func (h *Handler) GetAllTables(c *gin.Context) {

	tables, err := h.Service.GetAllTables()

	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"tables": tables,
	})
}
