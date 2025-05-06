package migration

import "gorm.io/gorm"

type v2 struct{}

func (v *v2) Migrate(db *gorm.DB) error {
	// Criar a tabela 'wear_amounts' se não existir
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS wear_amounts (
		    id UUID PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE,
			deleted_at TIMESTAMP WITH TIME ZONE,
			description VARCHAR(255) NOT NULL
		);
	`
	if err := db.Exec(createTableSQL).Error; err != nil {
		return err
	}

	// Adicionar coluna 'wear_amount' com referência à tabela 'wear_amounts'
	addColumns := `
		ALTER TABLE skins
		ADD COLUMN IF NOT EXISTS wear_amount UUID REFERENCES wear_amounts(id);
	`
	if err := db.Exec(addColumns).Error; err != nil {
		return err
	}

	// Excluir coluna 'drop_chance' da tabela 'skin', se existir
	dropColumnSQL := `ALTER TABLE skins DROP COLUMN IF EXISTS drop_chance;`
	if err := db.Exec(dropColumnSQL).Error; err != nil {
		return err
	}

	return nil
}
