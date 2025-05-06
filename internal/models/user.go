package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
)

type User struct {
	Base
	Name                    string         `json:"name" gorm:"not null"`
	TradeLink               string         `json:"tradeLink" gorm:"not null"`
	WalletID                uuid.UUID      `json:"walletID"`
	InventoryID             uuid.UUID      `json:"inventoryID"`
	SteamID                 string         `json:"steamId" gorm:"column:steam_id;uniqueIndex"`
	AvatarUrl               string         `json:"avatarUrl" gorm:"column:avatar_url"`
	PagarmeCustomerID       string         `json:"pagarmeCustomerID"`
	PagarmeCustomerMetadata datatypes.JSON `json:"pagarmeCustomerMetadata"`

	Wallet    *Wallet      `json:"wallet" gorm:"foreignKey:WalletID"`
	Inventory *Inventory   `json:"inventory" gorm:"foreignKey:InventoryID"`
	Address   *UserAddress `json:"address" gorm:"foreignKey:UserID;references:ID"`

	Provider string `json:"provider" gorm:"-"`
}

type UserAddress struct {
	Base
	UserID     uuid.UUID `json:"userId" gorm:"not null;index"`
	Number     string    `json:"number"`
	Address    string    `json:"address,omitempty"`
	Complement string    `json:"complement,omitempty"`
	ZipCode    string    `json:"zipCode,omitempty"`
	City       string    `json:"city,omitempty"`
	State      string    `json:"state,omitempty"`
	Country    string    `json:"country,omitempty"`
}
