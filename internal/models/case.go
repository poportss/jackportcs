package models

import (
	"github.com/gofrs/uuid"
)

type Case struct {
	Base
	Name    string      `json:"name"`
	Price   int64       `json:"price"`
	SkinsID []uuid.UUID `json:"skinsID"`
	Active  bool        `json:"active"`
	Skins   []*Skin     `gorm:"foreignKey:SkinsID"`
}
