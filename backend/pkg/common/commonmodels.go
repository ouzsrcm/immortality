package common

import "gorm.io/gorm"

type Bank struct {
	gorm.Model

	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Website     string `json:"website"`
	Description string `json:"description"`
}

type Branch struct {
	gorm.Model

	BankId      uint   `json:"bankId"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Website     string `json:"website"`
	Description string `json:"description"`
}

type CreditCardType struct {
	gorm.Model

	Name        string `json:"name"` // MasterCard, Visa, American Express, Discover
	Description string `json:"description"`
}
