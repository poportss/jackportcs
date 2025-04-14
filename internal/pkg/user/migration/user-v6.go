package migration

import "gorm.io/gorm"

type v6 struct{}

func (v *v6) Migrate(db *gorm.DB) error {
	migration := `
		CREATE TABLE IF NOT EXISTS user_addresses (
		 id UUID PRIMARY KEY,
    	 created_at TIMESTAMP WITH TIME ZONE,
    	 updated_at TIMESTAMP WITH TIME ZONE,
    	 deleted_at TIMESTAMP WITH TIME ZONE,

			user_id UUID NOT NULL,
			address VARCHAR(255),
			complement VARCHAR(255),
			zip_code VARCHAR(20),
			city VARCHAR(100),
			state VARCHAR(100),
			country VARCHAR(100),

			CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			CONSTRAINT uq_user_address UNIQUE (user_id)
		);`

	if err := db.Exec(migration).Error; err != nil {
		return err
	}

	return nil
}
