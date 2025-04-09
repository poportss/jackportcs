package migration

import "gorm.io/gorm"

type v6 struct{}

func (v *v6) Migrate(db *gorm.DB) error {
	alterColumnSQL := `ALTER TABLE payment_orders DROP COLUMN reference_id;`

	if err := db.Exec(alterColumnSQL).Error; err != nil {
		return err
	}

	addColumns := `ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS reference_id varchar(80)`

	if err := db.Exec(addColumns).Error; err != nil {
		return err
	}

	return nil
}
