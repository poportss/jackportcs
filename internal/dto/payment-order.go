package dto

import (
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/models"
	"time"
)

type CreateOrderRequest struct {
	Amount        int                  `json:"amount"`
	PaymentMethod models.PaymentMethod `json:"payment_method"`
	CardID        *uuid.UUID           `json:"cardID"`
	Card          *models.Card         `json:"card,omitempty"`
}

type CreateOrderResponse struct {
	ID            uuid.UUID               `json:"id"`
	CreatedAt     time.Time               `json:"createdAt"`
	Amount        int                     `json:"amount,omitempty" `
	Status        string                  `json:"status,omitempty" `
	PaymentMethod string                  `json:"paymentMethod,omitempty" `
	PIX           *models.PagarmeOrderPIX `json:"pix,omitempty"`
}
