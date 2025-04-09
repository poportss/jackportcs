package migration

import "gorm.io/gorm"

type v2 struct{}

func (v *v2) Migrate(db *gorm.DB) error {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS cards (
		id UUID PRIMARY KEY,
		created_at TIMESTAMP WITH TIME ZONE,
		updated_at TIMESTAMP WITH TIME ZONE,
		deleted_at TIMESTAMP WITH TIME ZONE,

		user_id UUID REFERENCES users(id),
		provider_card_id TEXT UNIQUE,
	    
		first_six_digits TEXT,
		last_four_digits TEXT,
		payment_method TEXT DEFAULT 'credit_card',

		holder_name TEXT NOT NULL,
		exp_month INTEGER NOT NULL,
		exp_year INTEGER NOT NULL,

		metadata JSONB,
		active BOOLEAN
	);`

	if err := db.Exec(createTableSQL).Error; err != nil {
		return err
	}

	createIndexes := `
		CREATE UNIQUE INDEX IF NOT EXISTS unq_provider_card ON cards(provider_card_id);
		CREATE INDEX IF NOT EXISTS idx_cards_user_id ON cards(user_id);
	`

	if err := db.Exec(createIndexes).Error; err != nil {
		return err
	}

	return nil
}
