package migration

import (
	"gorm.io/gorm"
)

type v1 struct{}

func (v *v1) Migrate(db *gorm.DB) error {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS schema_versions (
		    id UUID PRIMARY KEY,
			created_at TIMESTAMP WITH TIME ZONE,
			updated_at TIMESTAMP WITH TIME ZONE,
			deleted_at TIMESTAMP WITH TIME ZONE,
			service VARCHAR(255) NOT NULL,
			version INT NOT NULL
		)`

	if err := db.Exec(createTableSQL).Error; err != nil {
		return err
	}

	return nil
}
