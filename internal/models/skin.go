package models

import "github.com/gofrs/uuid"

type Skin struct {
	Base
	Name       string    `json:"name"`
	ImageURL   string    `json:"imageUrl"`
	Value      int64     `json:"value"`
	WearAmount uuid.UUID `json:"wearAmount"`
}

type WearAmount struct {
	Base
	Description string `json:"description"`
}
