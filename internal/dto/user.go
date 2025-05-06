package dto

import (
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/models"
)

type UserTradeLink struct {
	TradeLink string `json:"tradeLink"`
}

type UserAddress struct {
	Address    string `json:"address"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	ZipCode    string `json:"zipCode"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
}

type UserResponse struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	TradeLink string         `json:"tradeLink"`
	AvatarUrl string         `json:"avatarUrl"`
	Wallet    *models.Wallet `json:"wallet"`
}
