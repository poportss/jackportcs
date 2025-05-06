package dto

import (
	"github.com/gofrs/uuid"
	"time"
)

type Case struct {
	Name     string   `json:"name"`
	Price    int64    `json:"price"`
	ImageURL string   `json:"imageUrl"`
	SkinsID  []string `json:"skins"`
}

type CaseResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	ImageURL  string    `json:"imageUrl"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
}

type CaseDetailsResponse struct {
	ID                uuid.UUID           `json:"id"`
	CreatedAt         time.Time           `json:"createdAt"`
	ImageURL          string              `json:"imageUrl"`
	Name              string              `json:"name"`
	Price             int64               `json:"price"`
	CaseProbabilities []CaseProbabilities `json:"caseProbabilities"`
}

type CaseProbabilities struct {
	Skin        SkinMetadataResponse `json:"skin"`
	Probability int64                `json:"probability"`
}

type SteamIDRequest struct {
	SteamID string `json:"steamID"`
}
