package order

import (
	"time"

	"gorm.io/gorm"
)

const (
	// BasketStatusPending is the status of the basket when it is created.
	BasketStatusPending = "pending"

	// BasketStatusPaid is the status of the basket when it is paid.
	BasketStatusPaid = "paid"

	// BasketStatusCancelled is the status of the basket when it is cancelled.
	BasketStatusCancelled = "cancelled"
)

type Basket struct {
	gorm.Model

	BenefactorID uint `json:"benefactor_id"` // TODO: add benefactor id

	Amount         float64    `json:"amount"`
	Total          float64    `json:"total"`
	ClientIP       string     `json:"client_ip"`
	ExtraInfo      string     `json:"extra_info"`
	Status         string     `json:"status"` // pending, paid, cancelled
	LastStatusDate *time.Time `json:"status_date"`
}

type BasketInfo struct {
	gorm.Model

	BasketID    uint   `json:"basket_id"`
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

type BasketItem struct {
	gorm.Model

	BasketID    uint `json:"basket_id"`
	DonationID  uint `json:"donation_id"`
	Quantity    uint `json:"quantity"`
	Amount      uint `json:"amount"`
	SubTotal    uint `json:"sub_total"`
	DonateForID uint `json:"donate_for_id"` // Başkası için yapılan bağışlarda kullanılacak.

}
