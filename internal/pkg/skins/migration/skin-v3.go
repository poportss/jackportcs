package migration

import "gorm.io/gorm"

type v3 struct{}

func (v *v3) Migrate(db *gorm.DB) error {
	// Renomear a coluna 'imagem_url' para 'image_url' na tabela 'skins'
	renameColumnSQL := `ALTER TABLE skins RENAME COLUMN imageurl TO image_url;`
	if err := db.Exec(renameColumnSQL).Error; err != nil {
		return err
	}

	return nil
}
