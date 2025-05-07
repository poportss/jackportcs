package migration

import "gorm.io/gorm"

type v5 struct{}

func (v *v5) Migrate(db *gorm.DB) error {
	// Renomear a coluna 'wear_amount' para 'wear_amount_id'
	if err := db.Exec(`ALTER TABLE skins RENAME COLUMN wear_amount TO wear_amount_id`).Error; err != nil {
		return err
	}

	// Renomear a coluna 'rarity_type' para 'rarity_type_id'
	if err := db.Exec(`ALTER TABLE skins RENAME COLUMN rarity_type TO rarity_type_id`).Error; err != nil {
		return err
	}

	return nil
}
