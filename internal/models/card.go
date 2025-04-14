package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
)

type Card struct {
	Base

	UserID         uuid.UUID `json:"userID,omitempty"`
	ProviderCardID string    `json:"providerCardId,omitempty" gorm:"uniqueIndex:unq_provider_card"`

	FirstSixDigits string         `json:"firstSixDigits,omitempty"`
	LastFourDigits string         `json:"lastFourDigits,omitempty"`
	PaymentMethod  PaymentMethod  `json:"paymentMethod,omitempty" gorm:"default:credit_card"`
	Number         string         `json:"number,omitempty" gorm:"-" validate:"required"`
	HolderName     string         `json:"holderName,omitempty" validate:"required"`
	HolderDocument string         `json:"holderDocument,omitempty" gorm:"-"`
	ExpMonth       int            `json:"expMonth,omitempty" validate:"required"`
	ExpYear        int            `json:"expYear,omitempty" validate:"required"`
	Metadata       datatypes.JSON `json:"metadata"`
	Active         bool           `json:"active"`
}
