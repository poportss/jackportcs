package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
)

type PaymentOrder struct {
	Base
	UserID        uuid.UUID     `json:"userID,omitempty" gorm:"column:user_id"`
	CustomerID    string        `json:"customerID,omitempty" gorm:"column:customer_id"`
	ReferenceID   string        `json:"referenceId,omitempty" gorm:"column:reference_id" validate:"required"`
	Amount        int           `json:"amount,omitempty" gorm:"column:amount" validate:"required"` // Valor em centavos
	Status        OrderStatus   `json:"status,omitempty" gorm:"column:status;index"`               // pending, paid, canceled etc
	PaymentMethod PaymentMethod `json:"paymentMethod,omitempty" gorm:"column:payment_method" validate:"required"`

	// Armazena metadados e respostas do provedor (como Pagar.me)
	MetadataRaw       datatypes.JSON `json:"-" gorm:"column:metadata_raw"`       // Custom business metadata
	ProviderResponses datatypes.JSON `json:"-" gorm:"column:provider_responses"` // Resposta original da Pagar.me (orders/charges)

	Card *Card            `json:"card,omitempty" gorm:"-"`
	PIX  *PagarmeOrderPIX `json:"pix,omitempty" gorm:"-"`
}
