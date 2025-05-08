package dto

import (
	"github.com/gofrs/uuid"
	"time"
)

type Skin struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	ImageURL   string `json:"imageUrl"`
	WearAmount string `json:"wearAmount"`
	RarityType string `json:"rarityType"`
	Value      int64  `json:"value"`
}

type SkinOpenedCase struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	ImageURL        string    `json:"imageUrl"`
	WearAmount      string    `json:"wearAmount"`
	RarityType      string    `json:"rarityType"`
	Value           int64     `json:"value"`
	InventorySkinID uuid.UUID `json:"inventorySkinID"`
}

type RarityType struct {
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}

type WearAmount struct {
	Description string `json:"description"`
}

type SkinResponse struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	Name       string    `json:"name"`
	ImageURL   string    `json:"imageUrl"`
	WearAmount uuid.UUID `json:"wearAmount"`
	Value      int64     `json:"value"`
}

type SkinMetadataResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"imageUrl"`
	Value     int64     `json:"value"`
	Rarity    string    `json:"rarity"`
}
