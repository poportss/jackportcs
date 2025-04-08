package models

type Wallet struct {
	Base
	Balance float64 `gorm:"not null;default:0"`
}
