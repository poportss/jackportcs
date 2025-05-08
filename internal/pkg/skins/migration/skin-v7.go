package migration

import "gorm.io/gorm"

type v7 struct{}

func (v *v7) Migrate(db *gorm.DB) error {

	// Criar a tabela 'wear_amounts' se não existir
	createTableSQL := `
	CREATE TABLE user_transactions (
   			id UUID PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE,
			deleted_at TIMESTAMP WITH TIME ZONE,
    		user_id UUID NOT NULL,
    		type VARCHAR(255) NOT NULL,
    		description TEXT,
    		amount INTEGER NOT NULL,
    		balance_after INTEGER NOT NULL,
    		
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
	`
	if err := db.Exec(createTableSQL).Error; err != nil {
		return err
	}

	return nil
}
