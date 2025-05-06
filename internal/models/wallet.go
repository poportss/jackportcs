package models

type Wallet struct {
	Base
	Balance float64 `json:"balance"  gorm:"not null;default:0"`
}
