package migration

import "gorm.io/gorm"

type v4 struct{}

func (v *v4) Migrate(db *gorm.DB) error {
	// Criar a tabela 'wear_amounts' se não existir
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS rarity_types (
		    id UUID PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE,
			deleted_at TIMESTAMP WITH TIME ZONE,
			description VARCHAR(255) NOT NULL,
		    priority INTEGER NOT NULL
		);
	`
	if err := db.Exec(createTableSQL).Error; err != nil {
		return err
	}

	// Adicionar coluna 'wear_amount' com referência à tabela 'wear_amounts'
	addColumns := `
		ALTER TABLE skins
		ADD COLUMN IF NOT EXISTS rarity_type UUID REFERENCES rarity_types(id);
	`
	if err := db.Exec(addColumns).Error; err != nil {
		return err
	}

	return nil
}
