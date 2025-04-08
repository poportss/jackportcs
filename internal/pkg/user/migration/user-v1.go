package migration

import "gorm.io/gorm"

type v1 struct{}

func (v *v1) Migrate(db *gorm.DB) error {
	createWalletSQL := `
CREATE TABLE IF NOT EXISTS wallet (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    balance INTEGER NOT NULL DEFAULT 0);`

	createInventorySQL := `
CREATE TABLE IF NOT EXISTS inventory (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE);`

	createInventorySkinsSQL := `
CREATE TABLE IF NOT EXISTS inventory_skins (
    inventory_id UUID NOT NULL REFERENCES inventory(id) ON DELETE CASCADE,
    skin_id UUID NOT NULL REFERENCES skins(id) ON DELETE CASCADE,
    PRIMARY KEY (inventory_id, skin_id));`

	createUsersSQL := `
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    name VARCHAR(255) NOT NULL,
    trade_link TEXT NOT NULL,
    wallet_id UUID NOT NULL,
    inventory_id UUID NOT NULL,
    CONSTRAINT fk_wallet FOREIGN KEY (wallet_id) REFERENCES wallet(id),
    CONSTRAINT fk_inventory FOREIGN KEY (inventory_id) REFERENCES inventory(id));`

	// Executando todas as migrations em ordem correta
	if err := db.Exec(createWalletSQL).Error; err != nil {
		return err
	}
	if err := db.Exec(createInventorySQL).Error; err != nil {
		return err
	}
	if err := db.Exec(createInventorySkinsSQL).Error; err != nil {
		return err
	}
	if err := db.Exec(createUsersSQL).Error; err != nil {
		return err
	}

	return nil
}
