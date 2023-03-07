package benefactors

import (
	"errors"
	"immortality/pkg/common"

	"gorm.io/gorm"
)

const BENEFACTORS = "benefactors"
const COMPANIES = "companies"
const BENEFACTOR_ADDRESSES = "benefactor_addresses"
const BENEFACTOR_CREDIT_CARDS = "benefactor_credit_cards"
const BENEFACTOR_BANK_ACCOUNTS = "benefactor_bank_accounts"

type BenefactorStore struct {
	common.StoreBase
}

type IBenefactorStore interface {
	GetBenefactorById(id uint) (*Benefactor, error)
	GetBenefactors() ([]Benefactor, error)
	GetBenefactorsByCompanyId(companyId uint) ([]Benefactor, error)
	GetBenefactorsByUserId(userId uint) ([]Benefactor, error)

	CreateBenefactor(benefactor *Benefactor) error
	UpdateBenefactor(benefactor *Benefactor) error
	DeleteBenefactor(id uint) error

	GetBenefactorCreditCardById(id uint) (*BenefactorCreditCard, error)
	GetBenefactorCreditCards() ([]BenefactorCreditCard, error)
	GetBenefactorCreditCardsByBenefactorId(benefactorId uint) ([]BenefactorCreditCard, error)
	GetBenefactorBankAccountsByBenefactorId(benefactorId uint) ([]BenefactorBankAccount, error)

	CreateBenefactorCreditCard(benefactorCreditCard *BenefactorCreditCard) error
	UpdateBenefactorCreditCard(benefactorCreditCard *BenefactorCreditCard) error
	DeleteBenefactorCreditCard(id uint) error

	GetBenefactorAddressById(id uint) (*BenefactorAddress, error)
	GetBenefactorAddresses() ([]BenefactorAddress, error)
	GetBenefactorAddressesByBenefactorId(benefactorId uint) ([]BenefactorAddress, error)

	CreateBenefactorAddress(benefactorAddress *BenefactorAddress) error
	UpdateBenefactorAddress(benefactorAddress *BenefactorAddress) error
	DeleteBenefactorAddress(id uint) error

	GetCompanyById(id uint) (*Company, error)
	GetCompanies() ([]Company, error)

	CreateCompany(company *Company) error
	UpdateCompany(company *Company) error
	DeleteCompany(id uint) error
}

func (s *BenefactorStore) GetBenefactorById(id uint) (*Benefactor, error) {
	var benefactor Benefactor
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(BENEFACTORS).First(&benefactor)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return &benefactor, nil
}

func (s *BenefactorStore) GetBenefactors() ([]Benefactor, error) {
	var benefactors []Benefactor
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTORS).Find(&benefactors)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return benefactors, nil
}

func (s *BenefactorStore) GetBenefactorsByCompanyId(companyId uint) ([]Benefactor, error) {
	var benefactors []Benefactor
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("company_id = ?", companyId).Table(BENEFACTORS).Find(&benefactors)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return benefactors, nil
}

func (s *BenefactorStore) GetBenefactorsByUserId(userId uint) ([]Benefactor, error) {
	var benefactors []Benefactor
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("user_id = ?", userId).Table(BENEFACTORS).Find(&benefactors)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return benefactors, nil
}

func (s *BenefactorStore) CreateBenefactor(benefactor *Benefactor) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTORS).Create(benefactor)
		return res.Error
	})
	return txres
}

func (s *BenefactorStore) UpdateBenefactor(benefactor *Benefactor) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTORS).Save(benefactor)
		return res.Error
	})
	return txres
}

func (s *BenefactorStore) DeleteBenefactor(id uint) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(BENEFACTORS).Delete(&Benefactor{})
		return res.Error
	})
	return txres
}

func (s *BenefactorStore) GetBenefactorCreditCardById(id uint) (*BenefactorCreditCard, error) {
	var benefactorCreditCard BenefactorCreditCard
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(BENEFACTOR_CREDIT_CARDS).First(&benefactorCreditCard)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return &benefactorCreditCard, nil
}

func (s *BenefactorStore) GetBenefactorCreditCards() ([]BenefactorCreditCard, error) {
	var benefactorCreditCards []BenefactorCreditCard
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTOR_CREDIT_CARDS).Find(&benefactorCreditCards)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return benefactorCreditCards, nil
}

func (s *BenefactorStore) GetBenefactorCreditCardsByBenefactorId(benefactorId uint) ([]BenefactorCreditCard, error) {
	var benefactorCreditCards []BenefactorCreditCard
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("benefactor_id = ?", benefactorId).Table(BENEFACTOR_CREDIT_CARDS).Find(&benefactorCreditCards)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return benefactorCreditCards, nil
}

func (s *BenefactorStore) CreateBenefactorCreditCard(benefactorCreditCard *BenefactorCreditCard) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTOR_CREDIT_CARDS).Create(benefactorCreditCard)
		return res.Error
	})
	return txres
}

func (s *BenefactorStore) UpdateBenefactorCreditCard(benefactorCreditCard *BenefactorCreditCard) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTOR_CREDIT_CARDS).Save(benefactorCreditCard)
		return res.Error
	})
	return txres
}

func (s *BenefactorStore) DeleteBenefactorCreditCard(id uint) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(BENEFACTOR_CREDIT_CARDS).Delete(&BenefactorCreditCard{})
		return res.Error
	})
	return txres
}

func (s *BenefactorStore) GetBenefactorBankAccountById(id uint) (*BenefactorBankAccount, error) {
	var benefactorBankAccount BenefactorBankAccount
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(BENEFACTOR_BANK_ACCOUNTS).First(&benefactorBankAccount)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return &benefactorBankAccount, nil
}

func (s *BenefactorStore) GetBenefactorBankAccounts() ([]BenefactorBankAccount, error) {
	var benefactorBankAccounts []BenefactorBankAccount
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTOR_BANK_ACCOUNTS).Find(&benefactorBankAccounts)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return benefactorBankAccounts, nil
}

func (s *BenefactorStore) GetBenefactorBankAccountsByBenefactorId(benefactorId uint) ([]BenefactorBankAccount, error) {
	var benefactorBankAccounts []BenefactorBankAccount
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("benefactor_id = ?", benefactorId).Table(BENEFACTOR_BANK_ACCOUNTS).Find(&benefactorBankAccounts)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return benefactorBankAccounts, nil
}

func (s *BenefactorStore) GetBenefactorAddressById(id uint) (*BenefactorAddress, error) {
	var benefactorAddress BenefactorAddress
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(BENEFACTOR_ADDRESSES).First(&benefactorAddress)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return &benefactorAddress, nil
}

func (s *BenefactorStore) GetBenefactorAddresses() ([]BenefactorAddress, error) {
	var benefactorAddresses []BenefactorAddress
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTOR_ADDRESSES).Find(&benefactorAddresses)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return benefactorAddresses, nil
}

func (s *BenefactorStore) GetBenefactorAddressesByBenefactorId(benefactorId uint) ([]BenefactorAddress, error) {
	var benefactorAddresses []BenefactorAddress
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("benefactor_id = ?", benefactorId).Table(BENEFACTOR_ADDRESSES).Find(&benefactorAddresses)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return benefactorAddresses, nil
}

func (s *BenefactorStore) CreateBenefactorAddress(benefactorAddress *BenefactorAddress) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTOR_ADDRESSES).Create(benefactorAddress)
		return res.Error
	})
	return txres
}

func (s *BenefactorStore) UpdateBenefactorAddress(benefactorAddress *BenefactorAddress) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(BENEFACTOR_ADDRESSES).Save(benefactorAddress)
		return res.Error
	})
	return txres
}

func (s *BenefactorStore) DeleteBenefactorAddress(id uint) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(BENEFACTOR_ADDRESSES).Delete(&BenefactorAddress{})
		return res.Error
	})
	return txres
}
