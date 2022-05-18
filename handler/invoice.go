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

func (h *Handler) addInvoice(userID string) (string, error) {

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

	err = h.Service.AddInvoice(invoice)

	return userOrderID, err
}

func (h *Handler) GetInvoice(c *gin.Context) {
	claims, err := helper.ValidateToken(c)
	if err != nil {
		panic(err)
	}
	userID := claims.ID

	orderID, err := h.addInvoice(userID)
	if err != nil {
		panic(err)
	}

	invoice, err := h.Service.GetInvoice(orderID)
	if err != nil {
		panic(err)
	}

	details, err := h.Service.GetMoreDetailed(orderID)
	if err != nil {
		panic(err)
	}

	var foodDetails []entities.DummyInvoice

	for i := 0; i < len(details); i++ {
		dummyInvoice := entities.DummyInvoice{
			F_name:       details[i].F_name,
			F_count:      details[i].F_count,
			F_price:      details[i].F_price,
			F_totalPrice: details[i].F_totalPrice,
		}
		foodDetails = append(foodDetails, dummyInvoice)
	}

	invoice = entities.Invoice{
		Invoice_id:   invoice.Invoice_id,
		Order_ID:     invoice.Order_ID,
		Pay_date:     invoice.Pay_date,
		Total_cost:   invoice.Total_cost,
		Table_number: invoice.Table_number,
		Food_detail:  foodDetails,
	}

	c.IndentedJSON(http.StatusOK, invoice)
}
