package donation

import (
	"time"

	"gorm.io/gorm"
)

type MediaType string

const (
	// MediaTypeImage is image media type
	Image MediaType = "image"
	Video MediaType = "video"
	Sound MediaType = "sound"
)

// / Donation is a donation
type Donation struct {
	gorm.Model

	DonationGroupID uint `json:"donation_group_id"` // FK to DonationGroup

	Name             string `json:"name" example:"Donation Name"`                           // Donation name
	Description      string `json:"description" example:"Donation Description"`             // Donation description
	ShortDescription string `json:"short_description" example:"Donation Short Description"` // Donation short description
	Ordering         int    `json:"ordering"`                                               // Ordering

	DonationAmounts []DonationAmount `json:"donation_amounts"` // Donation amounts
}

// / DonationAmount is a donation amount
type DonationAmount struct {
	gorm.Model

	DonationID     uint      `json:"donation_id"`             // FK to Donation
	Amount         float64   `json:"amount" example:"100.00"` // Amount in TRY
	ValidStartDate time.Time `json:"valid_start_date"`        // Validity start date
	ValidEndDate   time.Time `json:"valid_end_date"`          // Validity end date
	IsActive       bool      `json:"is_active"`               // Is active
	IsDefault      bool      `json:"is_default"`              // Is default
}

// / DonationGroup is a group of donations
type DonationGroup struct {
	gorm.Model

	ParentId         uint   `json:"parent_id"`                                                    // FK to DonationGroup
	Name             string `json:"name" example:"Donation Group Name"`                           // Donation group name
	Description      string `json:"description" example:"Donation Group Description"`             // Donation group description
	ShortDescription string `json:"short_description" example:"Donation Short Group Description"` // Donation group short description
	Ordering         int    `json:"ordering"`                                                     // Ordering
}

// / DonationMedia is a media file for donation
type DonationMedia struct {
	gorm.Model

	DonationID uint `json:"donation_id"` // FK to Donation
	MediaID    uint `json:"media_id"`    // FK to Media
}

// / Media is a media file
type Media struct {
	gorm.Model

	Url       string `json:"url" example:"https://example.com/image.jpg"`
	Path      string `json:"path" example:"/path/to/image.jpg"` // Path to media file
	MediaType string `json:"media_type" example:"image"`        // image, video
	AutoPlay  bool   `json:"auto_play"`                         // Auto play media file
	Loop      bool   `json:"loop"`                              // Loop media file

}
