package migration

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

func SeedInitialUser(db *gorm.DB) error {
	// Gera os IDs
	walletID := uuid.Must(uuid.NewV4())
	inventoryID := uuid.Must(uuid.NewV4())
	userID := uuid.Must(uuid.NewV4())

	now := time.Now()

	// Insere a carteira
	if err := db.Exec(`
		INSERT INTO wallet (id, created_at, updated_at, balance)
		VALUES (?, ?, ?, ?)`,
		walletID, now, now, 1000).Error; err != nil {
		return err
	}

	// Insere o inventário
	if err := db.Exec(`
		INSERT INTO inventory (id, created_at, updated_at)
		VALUES (?, ?, ?)`,
		inventoryID, now, now).Error; err != nil {
		return err
	}

	// Insere o usuário
	if err := db.Exec(`
		INSERT INTO users (id, created_at, updated_at, name, trade_link, wallet_id, inventory_id)
		VALUES (?, ?, ?, ?, ?, ?, ?)`,
		userID, now, now, "JoãoCS", "https://steamcommunity.com/tradeoffer/new/?partner=123", walletID, inventoryID).Error; err != nil {
		return err
	}

	return nil
}
