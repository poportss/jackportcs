package migration

import (
	"gorm.io/gorm"
)

type v4 struct{}

func (v *v4) Migrate(db *gorm.DB) error {
	removeUniqueColumns := `
		ALTER TABLE inventory_skins DROP CONSTRAINT IF EXISTS inventory_skins_pkey
	`
	if err := db.Exec(removeUniqueColumns).Error; err != nil {
		return err
	}

	return nil
}
