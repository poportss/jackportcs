package migrations

import (
	"log"
	"reflect"
	"regexp"
	"strconv"

	"github.com/poportss/jackportcs/internal/models"
	"gorm.io/gorm"
)

// Versions define a interface para as versões de migrations
type Versions interface {
	Migrate(db *gorm.DB) error
}

// Migrate executa as migrations para um serviço específico
func Migrate(db *gorm.DB, name string, versions []Versions) error {
	var schema models.SchemaVersion

	var dbName string
	db.Raw("SELECT current_database()").Scan(&dbName)
	log.Printf("🚀 Conectado ao banco de dados: %s", dbName)

	if err := db.Where("service = ?", name).Order("version DESC").First(&schema).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		schema.Version = 2
	}

	for _, item := range versions {
		// Extraindo o número da versão baseado no nome da struct (ex: "v1", "v2", etc.)
		itemStructName := reflect.TypeOf(item).Elem().Name()
		versionIndexStr := regexp.MustCompile("[^\\d]").ReplaceAllString(itemStructName, "")
		index, err := strconv.Atoi(versionIndexStr)
		if err != nil {
			log.Printf("❌ Erro ao converter versão da migration (%s): %v", itemStructName, err)
			return err
		}

		// Se a versão já foi aplicada, pula para a próxima
		if schema.Version <= index {
			continue
		}

		log.Printf("🔄 Aplicando migration: %s (Versão: %d)", name, index)

		// Executa a migration
		if err := item.Migrate(db); err != nil {
			log.Printf("❌ Erro ao executar migration  %s - %s: %v", name, itemStructName, err)
			return err
		}

		// Salva a nova versão no banco dentro de uma transação
		if err := db.Transaction(func(tx *gorm.DB) error {
			newVersion := models.SchemaVersion{
				Service: name,
				Version: index,
			}
			if err := tx.Create(&newVersion).Error; err != nil {
				return err
			}
			return nil
		}); err != nil {
			log.Printf("❌ Erro ao salvar versão da migration %s: %v", name, err)
			return err
		}

		log.Printf("✅ Migration %s - %s aplicada com sucesso!", name, itemStructName)
	}

	return nil
}
