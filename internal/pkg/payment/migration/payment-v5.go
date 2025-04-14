package migration

import "gorm.io/gorm"

type v5 struct{}

func (v *v5) Migrate(db *gorm.DB) error {
	alterColumnSQL := `ALTER TABLE payment_orders  ALTER COLUMN reference_id TYPE VARCHAR(80);`

	if err := db.Exec(alterColumnSQL).Error; err != nil {
		return err
	}

	return nil

}
