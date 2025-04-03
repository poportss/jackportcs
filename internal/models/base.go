package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

// Base define um modelo base com UUID como chave primária
type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate define um UUID automaticamente antes da criação do registro
func (base *Base) BeforeCreate(_ *gorm.DB) error {
	if base.ID == uuid.Nil {
		nonce, err := uuid.NewV4()
		if err != nil {
			return err
		}
		base.ID = nonce
	}
	return nil
}
