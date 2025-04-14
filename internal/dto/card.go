package dto

import (
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/models"
	"time"
)

type Card struct {
	FirstSixDigits string               `json:"firstSixDigits"`
	LastFourDigits string               `json:"lastFourDigits"`
	PaymentMethod  models.PaymentMethod `json:"paymentMethod"`
	Number         string               `json:"number"  validate:"required"`
	HolderName     string               `json:"holderName" validate:"required"`
	HolderDocument string               `json:"holderDocument"`
	ExpMonth       int                  `json:"expMonth" validate:"required"`
	ExpYear        int                  `json:"expYear" validate:"required"`
	CVV            string               `json:"cvv" validate:"required,max=4,min=3"`
}

type CardResponse struct {
	ID             uuid.UUID            `json:"id"`
	CreatedAt      time.Time            `json:"createdAt"`
	LastFourDigits string               `json:"lastFourDigits"`
	PaymentMethod  models.PaymentMethod `json:"paymentMethod" `
	HolderName     string               `json:"holderName" `
	HolderDocument string               `json:"holderDocument"`
	ExpMonth       int                  `json:"expMonth"`
	ExpYear        int                  `json:"expYear"`
}
