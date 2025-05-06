package dto

import (
	"github.com/gofrs/uuid"
	"time"
)

type Skin struct {
	Name       string `json:"name"`
	ImageURL   string `json:"imageUrl"`
	WearAmount string `json:"wearAmount"`
	Value      int64  `json:"value"`
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
}
