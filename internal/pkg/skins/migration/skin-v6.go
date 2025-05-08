package migration

import "gorm.io/gorm"

type v6 struct{}

func (v *v6) Migrate(db *gorm.DB) error {

	// Adicionar coluna 'wear_amount' com referência à tabela 'wear_amounts'
	addColumns := `
		ALTER TABLE inventory_skins
		ADD COLUMN IF NOT EXISTS sold BOOLEAN DEFAULT FALSE;
	`
	if err := db.Exec(addColumns).Error; err != nil {
		return err
	}

	return nil
}
