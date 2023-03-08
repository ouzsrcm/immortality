package order

import (
	"time"

	"gorm.io/gorm"
)

const (
	// OrderStatusPending is the status of the order when it is created.
	OrderStatusPending = "pending"

	// OrderStatusPaid is the status of the basket when it is paid.
	OrderStatusPaid = "paid"

	// OrderStatusCancelled is the status of the basket when it is cancelled.
	OrderStatusCancelled = "cancelled"
)

type Order struct {
	gorm.Model

	DonationID     uint       `json:"donation_id"`
	BenefactorID   uint       `json:"benefactor_id"` // TODO: add benefactor id
	Amount         float64    `json:"amount"`
	Total          float64    `json:"total"`
	ClientIP       string     `json:"client_ip"`
	ExtraInfo      string     `json:"extra_info"`
	Status         string     `json:"status"` // pending, paid, cancelled
	LastStatusDate *time.Time `json:"status_date"`
}

type OrderInfo struct {
	gorm.Model

	OrderID     uint   `json:"order_id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	DistrictID  uint   `json:"district_id"`
	CityID      uint   `json:"city_id"`
	CountryID   uint   `json:"country_id"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type OrderItem struct {
	gorm.Model

	OrderID     uint `json:"order_id"`
	DonationID  uint `json:"donation_id"`
	Quantity    uint `json:"quantity"`
	Amount      uint `json:"amount"`
	SubTotal    uint `json:"sub_total"`
	DonateForID uint `json:"donate_for_id"` // Başkası için yapılan bağışlarda kullanılacak.
}
