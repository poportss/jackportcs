package models

import "github.com/gofrs/uuid"

type InventorySkins struct {
	Base
	InventoryID uuid.UUID `json:"inventoryID"`
	SkinID      uuid.UUID `json:"skinID"`
}
