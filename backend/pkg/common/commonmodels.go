package common

import "gorm.io/gorm"

type Bank struct {
	gorm.Model

	Name        string `json:"name" example:"Bank of America"`
	Address     string `json:"address" example:"1234 Main St`
	Phone       string `json:"phone" example:"555-555-5555"`
	Email       string `json:"email" example:"john.doe@gmail.com"`
	Website     string `json:"website" example:"https://www.google.com"`
	Description string `json:"description" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."`
}

type Branch struct {
	gorm.Model

	BankId      uint   `json:"bankId" example:"1"` // Banka bilgileri
	Name        string `json:"name" example:"Bank of America - Main Branch"`
	Address     string `json:"address" example:"1234 Main St`
	Phone       string `json:"phone" example:"555-555-5555"`
	Email       string `json:"email" example:"john.doe@gmail.com"`
	Website     string `json:"website" example:"https://www.google.com"`
	Description string `json:"description" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."`
}

type CreditCardType struct {
	gorm.Model

	Name        string `json:"name" example:"1"` // MasterCard, Visa, American Express, Discover
	Description string `json:"description" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."`
}
