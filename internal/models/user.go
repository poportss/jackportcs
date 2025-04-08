package models

import (
	"github.com/gofrs/uuid"
)

type User struct {
	Base
	Name        string    ` json:"name,omitempty" gorm:"not null"`
	TradeLink   string    ` json:"tradeLink,omitempty" gorm:"not null"`
	Password    string    ` json:"password,omitempty" gorm:"not null"`
	WalletID    uuid.UUID `json:"walletID,omitempty"`
	InventoryID uuid.UUID `json:"inventoryID,omitempty"`

	Wallet    Wallet    `json:"wallet" gorm:"foreignKey:WalletID" `
	Inventory Inventory ` json:"inventory"gorm:"foreignKey:InventoryID" `
}
