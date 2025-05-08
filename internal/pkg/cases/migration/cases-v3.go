package migration

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type InventorySkin struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type v3 struct{}

func (v *v3) Migrate(db *gorm.DB) error {
	// AutoMigrate adiciona colunas que não existem sem remover dados
	return db.AutoMigrate(&InventorySkin{})
}
