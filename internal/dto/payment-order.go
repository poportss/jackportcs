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

type WebhookOrder struct {
	ID        string    `json:"id"`
	Account   Account   `json:"account"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	Data      OrderData `json:"data"`
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OrderData struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Amount    int       `json:"amount"`
	Currency  string    `json:"currency"`
	Closed    bool      `json:"closed"`
	Status    string    `json:"status"`
	Items     []Item    `json:"items"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ClosedAt  time.Time `json:"closed_at"`
	Metadata  any       `json:"metadata"`
}

type Item struct {
	ID          string    `json:"id"`
	Amount      int       `json:"amount"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Status      string    `json:"status"`
	UpdatedAt   time.Time `json:"updated_at"`
}
