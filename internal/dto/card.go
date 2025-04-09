package dto

import "github.com/poportss/jackportcs/internal/models"

type Card struct {
	FirstSixDigits string               `json:"firstSixDigits,omitempty"`
	LastFourDigits string               `json:"lastFourDigits,omitempty"`
	PaymentMethod  models.PaymentMethod `json:"paymentMethod,omitempty" gorm:"default:credit_card"`
	Number         string               `json:"number,omitempty" gorm:"-" validate:"required"`
	HolderName     string               `json:"holderName,omitempty" validate:"required"`
	HolderDocument string               `json:"holderDocument,omitempty" gorm:"-"`
	ExpMonth       int                  `json:"expMonth,omitempty" validate:"required"`
	ExpYear        int                  `json:"expYear,omitempty" validate:"required"`
	CVV            string               `json:"cvv,omitempty" gorm:"-" validate:"required,max=4,min=3"`
}
