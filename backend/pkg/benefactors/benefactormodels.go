package benefactors

import "gorm.io/gorm"

type Benefactor struct {
	gorm.Model

	// UserId kullanıcı audit ve auth işlemlerinde kullanılacak
	UserId uint `json:"userId"`

	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CompanyId uint   `json:"company"`
	Position  string `json:"position"`
	Website   string `json:"website"`
	Linkedin  string `json:"linkedin"`
	Notes     string `json:"notes"`
}

type BenefactorAddress struct {
	gorm.Model

	BenefactorId uint   `json:"benefactorId"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Country      string `json:"country"`
	ZipCode      string `json:"zipCode"`
	Phone        string `json:"phone"`
}

type BenefactorCreditCard struct {
	gorm.Model

	BenefactorId     uint   `json:"benefactorId"`
	CreditCardTypeId uint   `json:"creditCardTypeId"` // MasterCard, Visa, American Express, Discover
	HolderName       string `json:"holderName"`
	CardNumber       string `json:"cardNumber"`
	ExpireDate       string `json:"expireDate"`
	CVV              string `json:"cvv"`
	IsActive         bool   `json:"isActive"`
}

type BenefactorBankAccount struct {
	gorm.Model

	BankId        uint   `json:"bankId"` // Banka bilgileri
	BenefactorId  uint   `json:"benefactorId"`
	Name          string `json:"name"`
	IBAN          string `json:"iban"`
	BIC           string `json:"bic"`
	AccountNumber string `json:"accountNumber"`
	IsActive      bool   `json:"isActive"`
}

type Company struct {
	gorm.Model

	Name      string `json:"name"`
	Address   string `json:"address"`
	TaxNumber string `json:"taxNumber"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Website   string `json:"website"`
}
