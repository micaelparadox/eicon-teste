package services

import (
	"errors"
	"test-eicon/models"
	"test-eicon/repositories"
)

type OrderService struct {
	repository *repositories.OrderRepository
}

func NewOrderService(repository *repositories.OrderRepository) *OrderService {
	return &OrderService{repository: repository}
}

func (service *OrderService) CreateOrder(order *models.Order) error {
	// Verifica se o número de controle já existe
	existingOrders, err := service.repository.GetOrdersByControlNumber(order.ControlNumber)
	if err != nil {
		return err
	}

	if len(existingOrders) > 0 {
		return errors.New("control number already exists")
	}

	return service.repository.CreateOrder(order)
}

func (service *OrderService) GetAllOrders() ([]models.Order, error) {
	return service.repository.GetAllOrders()
}

func (service *OrderService) GetOrdersByControlNumber(controlNumber string) ([]models.Order, error) {
	return service.repository.GetOrdersByControlNumber(controlNumber)
}

func (service *OrderService) GetOrdersByDate(date string) ([]models.Order, error) {
	return service.repository.GetOrdersByDate(date)
}
