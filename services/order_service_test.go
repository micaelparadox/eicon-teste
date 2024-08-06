package services

import (
	"test-eicon/models"
	"test-eicon/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Order{})
	return db
}

func TestCreateOrder(t *testing.T) {
	db := setupTestDB()
	repo := repositories.NewOrderRepository(db)
	service := NewOrderService(repo)

	order := &models.Order{
		ControlNumber: "999999",
		Name:          "Test Product",
		UnitPrice:     100,
		Quantity:      2,
		CustomerCode:  1,
	}

	err := service.CreateOrder(order)
	assert.Nil(t, err)

	var createdOrder models.Order
	db.First(&createdOrder, "control_number = ?", "999999")

	assert.Equal(t, "999999", createdOrder.ControlNumber)
	assert.Equal(t, "Test Product", createdOrder.Name)
	assert.Equal(t, 100.0, createdOrder.UnitPrice)
	assert.Equal(t, 2, createdOrder.Quantity)
	assert.Equal(t, 1, createdOrder.CustomerCode)
	assert.Equal(t, 200.0, createdOrder.TotalValue)
}
