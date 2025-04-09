package migration

import "gorm.io/gorm"

type v1 struct{}

func (v *v1) Migrate(db *gorm.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS payment_orders (
		id UUID PRIMARY KEY,
		created_at TIMESTAMP WITH TIME ZONE,
		updated_at TIMESTAMP WITH TIME ZONE,
		deleted_at TIMESTAMP WITH TIME ZONE,
		user_id UUID REFERENCES users(id),
		reference_id UUID UNIQUE NOT NULL,
		amount INTEGER NOT NULL,
		status varchar(80) NOT NULL,
		payment_method varchar(80) NOT NULL,
		metadata_raw JSONB,
		provider_responses JSONB,
		boleto_data JSONB
	);`

	if err := db.Exec(createTableSQL).Error; err != nil {
		return err
	}

	createIndexSQL := `
		CREATE INDEX IF NOT EXISTS idx_payment_orders_status ON payment_orders(status);
		CREATE INDEX IF NOT EXISTS idx_payment_orders_user_id ON payment_orders(user_id);
		CREATE INDEX IF NOT EXISTS idx_payment_orders_id ON payment_orders(id);
	`

	if err := db.Exec(createIndexSQL).Error; err != nil {
		return err
	}

	return nil
}
