package donation

import (
	"immortality/pkg/common"

	"gorm.io/gorm"
)

type DonationStore struct {
	common.StoreBase
}

func NewDonationStore() *DonationStore {
	store := new(DonationStore)
	store.Connect()
	return store
}

type IDonationService interface {
	GetDonationGroups() ([]DonationGroup, error)
	GetDonationGroup(id uint) (DonationGroup, error)
	CreateDonationGroup(donationGroup DonationGroup) (DonationGroup, error)
	UpdateDonationGroup(id uint, donationGroup DonationGroup) (DonationGroup, error)
	DeleteDonationGroup(id uint) error

	GetDonationsByDonationGroupId(groupId uint) ([]Donation, error)
	GetDonations() ([]Donation, error)
	GetDonation(id uint) (Donation, error)
	CreateDonation(donation Donation) (Donation, error)
	UpdateDonation(id uint, donation Donation) (Donation, error)
	DeleteDonation(id uint) error

	GetDonationAmounts() ([]DonationAmount, error)
	GetDonationAmount(id uint) (DonationAmount, error)
	CreateDonationAmount(donationAmount DonationAmount) (DonationAmount, error)
	UpdateDonationAmount(id uint, donationAmount DonationAmount) (DonationAmount, error)
	DeleteDonationAmount(id uint) error

	GetDonationMediasByDonationId(donationId uint) ([]DonationMedia, error)

	GetDonationMedias() ([]DonationMedia, error)
	GetDonationMedia(id uint) (DonationMedia, error)
	CreateDonationMedia(donationMedia DonationMedia) (DonationMedia, error)
	UpdateDonationMedia(id uint, donationMedia DonationMedia) (DonationMedia, error)
	DeleteDonationMedia(id uint) error
}

func (s *DonationStore) GetDonationGroups() (*[]DonationGroup, error) {
	var donationGroups []DonationGroup
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Find(&donationGroups)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationGroups, nil
}

func (s *DonationStore) GetDonationGroup(id uint) (*DonationGroup, error) {
	var donationGroup DonationGroup
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.First(&donationGroup, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationGroup, nil
}

func (s *DonationStore) CreateDonationGroup(donationGroup DonationGroup) (*DonationGroup, error) {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&donationGroup)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationGroup, nil
}

func (s *DonationStore) UpdateDonationGroup(id uint, donationGroup DonationGroup) (*DonationGroup, error) {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&donationGroup).Where("id = ?", id).Updates(donationGroup)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationGroup, nil
}

func (s *DonationStore) DeleteDonationGroup(id uint) error {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Delete(&DonationGroup{}, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return tempres
	}
	return nil
}

func (s *DonationStore) GetDonation(id uint) (*Donation, error) {
	var donation Donation
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.First(&donation, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donation, nil
}

func (s *DonationStore) GetDonations() (*[]Donation, error) {
	var donations []Donation
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Find(&donations)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donations, nil
}

func (s *DonationStore) CreateDonation(donation Donation) (*Donation, error) {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&donation)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donation, nil
}

func (s *DonationStore) UpdateDonation(id uint, donation Donation) (*Donation, error) {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&donation).Where("id = ?", id).Updates(donation)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donation, nil
}

func (s *DonationStore) DeleteDonation(id uint) error {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Delete(&Donation{}, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return tempres
	}
	return nil
}

func (s *DonationStore) GetDonationAmount(id uint) (*DonationAmount, error) {
	var donationAmount DonationAmount
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.First(&donationAmount, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationAmount, nil
}

func (s *DonationStore) GetDonationAmounts() (*[]DonationAmount, error) {
	var donationAmounts []DonationAmount
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Find(&donationAmounts)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationAmounts, nil
}

func (s *DonationStore) CreateDonationAmount(donationAmount DonationAmount) (*DonationAmount, error) {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&donationAmount)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationAmount, nil
}

func (s *DonationStore) UpdateDonationAmount(id uint, donationAmount DonationAmount) (*DonationAmount, error) {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&donationAmount).Where("id = ?", id).Updates(donationAmount)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationAmount, nil
}

func (s *DonationStore) DeleteDonationAmount(id uint) error {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Delete(&DonationAmount{}, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return tempres
	}
	return nil
}

func (s *DonationStore) GetDonationMedias() (*[]DonationMedia, error) {
	var donationMedias []DonationMedia
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Find(&donationMedias)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationMedias, nil
}

func (s *DonationStore) GetDonationMedia(id uint) (*DonationMedia, error) {
	var donationMedia DonationMedia
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.First(&donationMedia, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationMedia, nil
}

func (s *DonationStore) CreateDonationMedia(donationMedia DonationMedia) (*DonationMedia, error) {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&donationMedia)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationMedia, nil
}

func (s *DonationStore) UpdateDonationMedia(id uint, donationMedia DonationMedia) (*DonationMedia, error) {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&donationMedia).Where("id = ?", id).Updates(donationMedia)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationMedia, nil
}

func (s *DonationStore) DeleteDonationMedia(id uint) error {
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Delete(&DonationMedia{}, id)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return tempres
	}
	return nil
}

func (s *DonationStore) GetDonationMediasByDonationId(donationId uint) (*[]DonationMedia, error) {
	var donationMedias []DonationMedia
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("donation_id = ?", donationId).Find(&donationMedias)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donationMedias, nil
}

func (s *DonationStore) GetDonationsByDonationGroupId(donationGroupId uint) (*[]Donation, error) {
	var donations []Donation
	tempres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("donation_group_id = ?", donationGroupId).Find(&donations)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if tempres != nil {
		return nil, tempres
	}
	return &donations, nil
}
