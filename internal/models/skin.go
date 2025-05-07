package models

import "github.com/gofrs/uuid"

type Skin struct {
	Base
	Name         string     `json:"name"`
	ImageURL     string     `json:"imageUrl"`
	Value        int64      `json:"value"`
	WearAmountID uuid.UUID  `json:"wearAmountID"`
	RarityTypeID uuid.UUID  `json:"rarityType"`
	RarityType   RarityType `gorm:"foreignKey:RarityTypeID"`
	WearAmount   WearAmount `gorm:"foreignKey:WearAmountID""`
}

type WearAmount struct {
	Base
	Description string `json:"description"`
}

type RarityType struct {
	Base
	Description string `json:"description"`
	Priority    int    `json:"priority"`
}
