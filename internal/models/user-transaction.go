package models

import (
	"github.com/gofrs/uuid"
)

type TransactionType string

const (
	TransactionTypeDeposit    TransactionType = "deposit"
	TransactionTypeBoxOpen    TransactionType = "box_open"
	TransactionTypeSkinSale   TransactionType = "skin_sale"
	TransactionTypeCaseOpen   TransactionType = "case_open"
	TransactionTypeSkinRefund TransactionType = "skin_refund"
)

type UserTransaction struct {
	Base
	UserID       uuid.UUID
	Type         TransactionType // "deposit", "box_open", "skin_sale", etc.
	Description  string
	Amount       int64 // Valor positivo ou negativo
	BalanceAfter int64 // Saldo do usuário após a transação
}
