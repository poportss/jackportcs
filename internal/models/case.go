package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
)

type Case struct {
	Base
	Name            string            `json:"name"`
	Price           int64             `json:"price"`
	Active          bool              `json:"active"`
	ImageURL        string            `json:"imageUrl"`
	Probabilities   datatypes.JSON    `gorm:"type:jsonb"`
	CaseProbability []CaseProbability `gorm:"-" json:"caseProbability"`
}

type CaseProbability struct {
	SkinID      uuid.UUID `json:"skinID"`
	Probability int64     `json:"probability"`
}
