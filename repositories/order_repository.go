package repositories

import (
	"test-eicon/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) CreateOrder(order *models.Order) error {
	return r.DB.Create(order).Error
}

func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) GetOrdersByControlNumber(controlNumber string) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Where("control_number = ?", controlNumber).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) GetOrdersByDate(date string) ([]models.Order, error) {
	var orders []models.Order
	err := r.DB.Where("DATE(registration_date) = ?", date).Find(&orders).Error
	return orders, err
}
