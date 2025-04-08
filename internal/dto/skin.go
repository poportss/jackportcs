package dto

import (
	"github.com/gofrs/uuid"
	"time"
)

type Skin struct {
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	DropChance int64  `json:"drop_chance"` // Probabilidade de drop (0.0 a 1.0)
	Value      int64  `json:"value"`
}

type SkinResponse struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	Name       string    `json:"name"`
	ImageURL   string    `json:"image_url"`
	DropChance int64     `json:"drop_chance"` // Probabilidade de drop (0.0 a 1.0)
	Value      int64     `json:"value"`
}
