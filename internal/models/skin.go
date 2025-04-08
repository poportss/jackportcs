package models

type Skin struct {
	Base
	Name       string `json:"name"`
	ImageURL   string `json:"image_url"`
	DropChance int64  `json:"drop_chance"` // Probabilidade de drop (0.0 a 1.0)
	Value      int64  `json:"value"`
}
