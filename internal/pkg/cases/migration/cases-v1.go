package migration

import (
	"gorm.io/gorm"
)

type v1 struct{}

func (v *v1) Migrate(db *gorm.DB) error {
	// Criar a tabela cases se não existir
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS cases (
		    id UUID PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE,
			deleted_at TIMESTAMP WITH TIME ZONE,
			name varchar(255) NOT NULL,
		    price integer NOT NULL,
		    skins_id uuid[] NOT NULL,
		    active boolean NOT NULL DEFAULT TRUE)`

	if err := db.Exec(createTableSQL).Error; err != nil {
		return err
	}

	return nil
}
