package migration

import "gorm.io/gorm"

type v1 struct{}

func (v *v1) Migrate(db *gorm.DB) error {
	// Criar a tabela cases se não existir
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS skins (
		    id UUID PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE,
			deleted_at TIMESTAMP WITH TIME ZONE,
			name varchar(255) NOT NULL,
			ImageURL text NOT NULL,
		    drop_chance integer NOT NULL,
		    value integer NOT NULL)`

	if err := db.Exec(createTableSQL).Error; err != nil {
		return err
	}

	return nil
}
