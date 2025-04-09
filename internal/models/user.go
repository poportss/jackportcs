package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
)

type User struct {
	Base
	Name                    string         ` json:"name" gorm:"not null"`
	TradeLink               string         ` json:"tradeLink" gorm:"not null"`
	WalletID                uuid.UUID      `json:"walletID"`
	InventoryID             uuid.UUID      `json:"inventoryID"`
	SteamID                 string         `json:"steamId" gorm:"column:steam_id;uniqueIndex"`
	AvatarUrl               string         `json:"avatarUrl" gorm:"column:avatar_url"`
	PagarmeCustomerID       string         `json:"pagarmeCustomerID"`
	PagarmeCustomerMetadata datatypes.JSON `json:"pagarmeCustomerMetadata"`

	Wallet    Wallet    `json:"wallet" gorm:"foreignKey:WalletID" `
	Inventory Inventory ` json:"inventory"gorm:"foreignKey:InventoryID" `
	Provider  string
}
