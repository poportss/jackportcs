package models

type Inventory struct {
	Base
	Skins []Skin `gorm:"many2many:inventory_skins"`
}
