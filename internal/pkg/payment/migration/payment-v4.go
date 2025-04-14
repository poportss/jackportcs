package migration

import "gorm.io/gorm"

type v4 struct{}

func (v *v4) Migrate(db *gorm.DB) error {
	addColumns := `ALTER TABLE payment_orders ADD COLUMN IF NOT EXISTS customer_id varchar(80) UNIQUE`

	if err := db.Exec(addColumns).Error; err != nil {
		return err
	}

	return nil

}
