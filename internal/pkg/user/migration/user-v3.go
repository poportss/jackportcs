package migration

import "gorm.io/gorm"

type v3 struct{}

func (v *v3) Migrate(db *gorm.DB) error {
	addColumns := `
	ALTER TABLE users
	ADD COLUMN IF NOT EXISTS pagarme_customer_id UUID,
	ADD COLUMN IF NOT EXISTS pagarme_customer_metadata JSONB;
	`

	if err := db.Exec(addColumns).Error; err != nil {
		return err
	}

	return nil
}
