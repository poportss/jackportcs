package migration

import "gorm.io/gorm"

type v7 struct{}

func (v *v7) Migrate(db *gorm.DB) error {
	addColumns := `ALTER TABLE user_addresses ADD COLUMN IF NOT EXISTS number varchar(30)`

	if err := db.Exec(addColumns).Error; err != nil {
		return err
	}

	return nil
}
