package migration

import "gorm.io/gorm"

type v5 struct{}

func (v *v5) Migrate(db *gorm.DB) error {
	migration := `
		ALTER TABLE users 
		    DROP COLUMN IF EXISTS avatarurl,
		    ADD COLUMN IF NOT EXISTS avatar_url TEXT;`

	if err := db.Exec(migration).Error; err != nil {
		return err
	}

	return nil
}
