package migration

import "gorm.io/gorm"

type v3 struct{}

func (v *v3) Migrate(db *gorm.DB) error {
	alterColumnSQL := `ALTER TABLE users  ALTER COLUMN pagarme_customer_id TYPE VARCHAR(80);`

	if err := db.Exec(alterColumnSQL).Error; err != nil {
		return err
	}

	return nil
}
