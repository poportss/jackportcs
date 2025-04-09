package migration

import "gorm.io/gorm"

type v4 struct{}

func (v *v4) Migrate(db *gorm.DB) error {
	addColumns := `
		ALTER TABLE users 
		    ADD COLUMN IF NOT EXISTS steam_id TEXT UNIQUE,
		    ADD COLUMN IF NOT EXISTS AvatarUrl TEXT;`

	if err := db.Exec(addColumns).Error; err != nil {
		return err
	}

	return nil
}
