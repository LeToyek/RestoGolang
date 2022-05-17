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

	var foodsName []string
	var foodsCount []int
	var foodsPrice []int
	var totalPrices []int

	for _, d := range details {
		foodsName = append(foodsName, d.F_name)
		foodsCount = append(foodsCount, d.F_count)
		foodsPrice = append(foodsPrice, d.F_price)
		totalPrices = append(totalPrices, d.F_totalPrice)
	}

	invoice = entities.Invoice{
		Invoice_id:   invoice.Invoice_id,
		Order_ID:     invoice.Order_ID,
		Pay_date:     invoice.Pay_date,
		Total_cost:   invoice.Total_cost,
		Table_number: invoice.Table_number,
		Food_count:   foodsCount,
		Food_price:   foodsPrice,
		Total_price:  totalPrices,
		Food_name:    foodsName,
	}

	c.IndentedJSON(http.StatusOK, invoice)
}
