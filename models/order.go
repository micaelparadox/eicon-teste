package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID               uint      `gorm:"primaryKey" json:"id" xml:"id"`
	ControlNumber    string    `gorm:"unique;not null" json:"control_number" xml:"control_number"`
	RegistrationDate time.Time `gorm:"not null" json:"registration_date" xml:"registration_date"`
	Name             string    `gorm:"not null" json:"name" xml:"name"`
	UnitPrice        float64   `gorm:"not null" json:"unit_price" xml:"unit_price"`
	Quantity         int       `gorm:"default:1" json:"quantity" xml:"quantity"`
	CustomerCode     int       `gorm:"not null" json:"customer_code" xml:"customer_code"`
	TotalValue       float64   `gorm:"not null" json:"total_value" xml:"total_value"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	if order.RegistrationDate.IsZero() {
		order.RegistrationDate = time.Now()
	}
	if order.Quantity == 0 {
		order.Quantity = 1
	}
	order.TotalValue = order.CalculateTotalValue()
	return
}

func (order *Order) CalculateTotalValue() float64 {
	total := float64(order.Quantity) * order.UnitPrice
	if order.Quantity >= 10 {
		total *= 0.9 // 10% de desconto
	} else if order.Quantity > 5 {
		total *= 0.95 // 5% de desconto
	}
	return total
}
