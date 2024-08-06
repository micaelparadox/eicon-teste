package controllers

import (
	"net/http"
	"test-eicon/models"
	"test-eicon/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var orders []models.Order

	contentType := c.GetHeader("Content-Type")
	if contentType == "application/xml" {
		if err := c.ShouldBindXML(&orders); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	} else {
		if err := c.ShouldBindJSON(&orders); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if len(orders) > 10 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Maximum 10 orders allowed"})
		return
	}

	for _, order := range orders {
		if order.ControlNumber == "" || order.Name == "" || order.UnitPrice == 0 || order.CustomerCode == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
			return
		}
		if err := ctrl.service.CreateOrder(&order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
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
