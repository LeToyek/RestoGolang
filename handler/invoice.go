package handler

import (
	"fmt"
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

	orderIDs, err := h.Service.GetOrderID(userID)
	if err != nil {
		panic(err)
	}

	userOrderID := orderIDs[len(orderIDs)-1]
	fmt.Println(userOrderID)

	prices, err := h.Service.GetPrice(userOrderID)
	if err != nil {
		panic(err)
	}

	var totalPrice int

	for _, p := range prices {
		totalPrice += p
	}

	fmt.Println(totalPrice)
	invoice := entities.Invoice{
		Invoice_id: nanoid.New(),
		Order_ID:   userOrderID,
		Pay_date:   time.Now().String(),
		Total_cost: float64(totalPrice),
	}

	h.Service.AddInvoice(invoice)
	c.IndentedJSON(http.StatusOK, gin.H{
		"created": invoice.Invoice_id,
	})
}

// func (h *Handler) GetInvoice(c *gin.Context) {
// 	claims, err := helper.ValidateToken(c)
// 	if err != nil {
// 		panic(err)
// 	}
// 	invoice,err := h.Service.GetInvoice(claims.ID)
// 	detail,err := h.Service.GetMoreDetailed(claims.ID)

// 	invoice = entities.Invoice{
// 		Table_number: detail.Table_number,
// 		Food_name: ,
// 	}

// 	c.IndentedJSON(http.StatusOK,i)
// }
