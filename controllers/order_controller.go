package controllers

import (
	"fmt"
	"net/http"
	"test-eicon/models"
	"test-eicon/services"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

type OrdersXMLWrapper struct {
	Orders []models.Order `xml:"order"`
}

func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var orders []models.Order

	contentType := c.GetHeader("Content-Type")
	if contentType == "application/xml" {
		var ordersWrapper OrdersXMLWrapper
		if err := c.ShouldBindXML(&ordersWrapper); err != nil {
			fmt.Printf("Error binding XML: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		orders = ordersWrapper.Orders
	} else {
		if err := c.ShouldBindJSON(&orders); err != nil {
			fmt.Printf("Error binding JSON: %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	fmt.Printf("Received Orders: %+v\n", orders)

	if len(orders) > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 10 orders allowed"})
		return
	}

	for i, order := range orders {
		if order.RegistrationDate.IsZero() {
			order.RegistrationDate = time.Now()
		}
		order.TotalValue = order.CalculateTotalValue()
		fmt.Printf("Inspect Order: %+v\n", order)
		if order.ControlNumber == "" || order.Name == "" || order.UnitPrice == 0 || order.CustomerCode == 0 {
			fmt.Printf("Missing Fields in Order: %+v\n", order)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}
		if err := ctrl.service.CreateOrder(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		orders[i] = order // Atualiza o pedido com os dados preenchidos
	}

	c.JSON(http.StatusOK, orders)
}

func (ctrl *OrderController) GetOrders(c *gin.Context) {
	controlNumber := c.Query("control_number")
	date := c.Query("date")

	var orders []models.Order
	var err error

	if controlNumber != "" {
		orders, err = ctrl.service.GetOrdersByControlNumber(controlNumber)
	} else if date != "" {
		orders, err = ctrl.service.GetOrdersByDate(date)
	} else {
		orders, err = ctrl.service.GetAllOrders()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}
