package models

type SchemaVersion struct {
	Base
	Service string `gorm:"service"`
	Version int    `gorm:"version"`
}
