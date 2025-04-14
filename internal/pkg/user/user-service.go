package user

import (
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
)

func (s *srv) createUserTradeLink(userTradeLink string, userId uuid.UUID) error {

	err := s.DB.Model(&models.User{}).Where("id = ?", userId).UpdateColumn("trade_link", userTradeLink).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *srv) createUserAddress(userAddressRequest dto.UserAddress, userId uuid.UUID) error {

	userAddress := &models.UserAddress{
		UserID:     userId,
		Address:    userAddressRequest.Address,
		Complement: userAddressRequest.Complement,
		Number:     userAddressRequest.Number,
		ZipCode:    userAddressRequest.ZipCode,
		City:       userAddressRequest.City,
		State:      userAddressRequest.State,
		Country:    userAddressRequest.Country,
	}

	if err := s.DB.Model(&models.UserAddress{}).Create(&userAddress).Error; err != nil {
		return err
	}

	return nil
}
