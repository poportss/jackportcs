package migration

import (
	"gorm.io/gorm"
)

type v2 struct{}

func (v *v2) Migrate(db *gorm.DB) error {
	addColumns := `
		ALTER TABLE cases ADD COLUMN IF NOT EXISTS probabilities jsonb;
		ALTER TABLE cases ADD COLUMN IF NOT EXISTS image_url text;
	`
	if err := db.Exec(addColumns).Error; err != nil {
		return err
	}

	dropColumnSQL := `ALTER TABLE cases DROP COLUMN IF EXISTS skins_id;`
	if err := db.Exec(dropColumnSQL).Error; err != nil {
		return err
	}

	return nil
}
