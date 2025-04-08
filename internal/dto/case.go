package dto

import (
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/models"
	"time"
)

type Case struct {
	Name    string      `json:"name"`
	Price   int64       `json:"price"`
	SkinsID []uuid.UUID `json:"skins"`
}

type CaseResponse struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	Name      string         `json:"name"`
	Price     int64          `json:"price"`
	Skins     []*models.Skin `json:"skins"`
}
