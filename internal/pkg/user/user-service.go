package user

import (
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/models"
)

func (s *srv) createUserTradeLink(userTradeLink string, userId uuid.UUID) error {

	err := s.DB.Model(&models.User{}).Where("id = ?", userId).UpdateColumn("trade_link", userTradeLink).Error
	if err != nil {
		return err
	}

	return nil
}
