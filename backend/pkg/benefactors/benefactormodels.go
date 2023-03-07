package benefactors

import "gorm.io/gorm"

type Benefactor struct {
	gorm.Model

	// UserId kullanıcı audit ve auth işlemlerinde kullanılacak
	UserId uint `json:"userId"`

	FirstName string `json:"firstName" example:"John"` // Kullanıcı adı
	LastName  string `json:"lastName" example:"Doe"`   // Soyadı
	Email     string `json:"email" example:"john.doe@gmail.com" `
	Phone     string `json:"phone" example:"555-555-5555"`
	CompanyId uint   `json:"company" example:"1"`                                                                                                                         // Şirket bilgileri
	Position  string `json:"position" example:"Software Engineer"`                                                                                                        // Pozisyon bilgileri
	Website   string `json:"website" 	example:"https://www.google.com"`                                                                                                   // Web sitesi bilgileri
	Linkedin  string `json:"linkedin" example:"https://www.linkedin.com"`                                                                                                 // Linkedin bilgileri
	Notes     string `json:"notes" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."` // Notlar
}

type BenefactorAddress struct {
	gorm.Model

	BenefactorId uint   `json:"benefactorId" example:"1"`
	Name         string `json:"name" example:"Home Address"`
	Address      string `json:"address" example:"1234 Main St`
	City         string `json:"city" example:"Istanbul"`
	Country      string `json:"country" example:"Turkey"`
	ZipCode      string `json:"zipCode" example:"34000"`
	Phone        string `json:"phone" example:"555-555-5555"`
}

type BenefactorCreditCard struct {
	gorm.Model

	BenefactorId     uint   `json:"benefactorId" example:"1"`
	CreditCardTypeId uint   `json:"creditCardTypeId" example:"1"` // MasterCard, Visa, American Express, Discover
	HolderName       string `json:"holderName" example:"John Doe"`
	CardNumber       string `json:"cardNumber" example:"1234 1234 1234 1234"` // 1234 1234 1234 1234
	ExpireDate       string `json:"expireDate" example:"12/2022"`             // 12/2022
	CVV              string `json:"cvv" example:"123"`                        // 123
	IsActive         bool   `json:"isActive" example:"true"`                  // true
}

type BenefactorBankAccount struct {
	gorm.Model

	BankId        uint   `json:"bankId" example:"1"` // Banka bilgileri
	BenefactorId  uint   `json:"benefactorId" example:"1"`
	Name          string `json:"name" example:"John Doe"`                      // Kullanıcı adı
	IBAN          string `json:"iban" example:"TR123456789012345678901234"`    // TR123456789012345678901234
	BIC           string `json:"bic" example:"TR123456789012345678901234"`     // TR123456789012345678901234
	AccountNumber string `json:"accountNumber" example:"12345678901234567890"` // 12345678901234567890
	IsActive      bool   `json:"isActive" example:"true"`                      // true
}

type Company struct {
	gorm.Model

	Name      string `json:"name" example:"Google Inc."`
	Address   string `json:"address" example:"1600 Amphitheatre Parkway"`
	TaxNumber string `json:"taxNumber" example:"1234567890"`
	Phone     string `json:"phone" example:"555-555-5555"`
	Email     string `json:"email" example:"john.doe@gmail.com""`
	Website   string `json:"website" example:"https://www.google.com"`
}
