package handler

import (
	"net/http"
	"resto/entities"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/gin-gonic/gin"

	helper "resto/helper"
)

func (h *Handler) AddInvoice(c *gin.Context) {
	claims, err := helper.ValidateToken(c)
	if err != nil {
		panic(err)
	}

	userID := claims.ID
	prices, err := h.Service.GetPrice(userID)
	if err != nil {
		panic(err)
	}
	var totalPrice int

	for _, p := range prices {
		totalPrice += p
	}

	invoice := entities.Invoice{
		Invoice_id: nanoid.New(),
		User_id:    userID,
		Pay_date:   time.Now().String(),
		Total_cost: float64(totalPrice),
	}

	h.Service.AddInvoice(invoice)
	c.IndentedJSON(http.StatusOK, gin.H{
		"created": invoice.Invoice_id,
	})
}

func (h *Handler) GetInvoice(c *gin.Context) {
	claims, err := helper.ValidateToken(c)
	if err != nil {
		panic(err)
	}
	invoice,err := h.Service.GetInvoice(claims.ID)
	detail,err := h.Service.GetMoreDetailed(claims.ID)

	invoice = entities.Invoice{
		Table_number: detail.Table_number,
		Food_name: ,
	}

	c.IndentedJSON(http.StatusOK,i)
}
