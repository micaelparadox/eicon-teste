package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID               uint      `gorm:"primaryKey"`
	ControlNumber    string    `gorm:"unique;not null"`
	RegistrationDate time.Time `gorm:"not null"`
	Name             string    `gorm:"not null"`
	UnitPrice        float64   `gorm:"not null"`
	Quantity         int       `gorm:"default:1"`
	CustomerCode     int       `gorm:"not null"`
	TotalValue       float64   `gorm:"not null"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if order.RegistrationDate.IsZero() {
		order.RegistrationDate = time.Now()
	}
	if order.Quantity == 0 {
		order.Quantity = 1
	}
	order.TotalValue = order.calculateTotalValue()
	return
}

func (order *Order) calculateTotalValue() float64 {
	total := float64(order.Quantity) * order.UnitPrice
	if order.Quantity >= 10 {
		total *= 0.9 // 10% de desconto
	} else if order.Quantity > 5 {
		total *= 0.95 // 5% de desconto
	}
	return total
}
