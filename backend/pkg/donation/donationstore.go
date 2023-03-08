package donation

import "immortality/pkg/common"

type DonationStore struct {
	common.StoreBase
}

type IDonationService interface {
	GetDonationGroups() ([]DonationGroup, error)
	GetDonationGroup(id uint) (DonationGroup, error)
	CreateDonationGroup(donationGroup DonationGroup) (DonationGroup, error)
	UpdateDonationGroup(id uint, donationGroup DonationGroup) (DonationGroup, error)
	DeleteDonationGroup(id uint) error

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

	GetDonationMedias() ([]DonationMedia, error)
	GetDonationMedia(id uint) (DonationMedia, error)
	CreateDonationMedia(donationMedia DonationMedia) (DonationMedia, error)
	UpdateDonationMedia(id uint, donationMedia DonationMedia) (DonationMedia, error)
	DeleteDonationMedia(id uint) error

	GetDonationGroupDonations(id uint) ([]Donation, error)
	GetDonationDonationAmounts(id uint) ([]DonationAmount, error)
	GetDonationDonationMedias(id uint) ([]DonationMedia, error)
}
