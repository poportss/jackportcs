package skins

import (
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
)

func (s *srv) createSkin(skin dto.Skin) (*models.Skin, error) {

	skinModel := &models.Skin{
		Name:         skin.Name,
		ImageURL:     skin.ImageURL,
		Value:        skin.Value,
		WearAmountID: uuid.FromStringOrNil(skin.WearAmount),
		RarityTypeID: uuid.FromStringOrNil(skin.RarityType),
	}

	err := s.DB.Create(&skinModel).Error
	if err != nil {
		return nil, err
	}

	return skinModel, nil
}

func (s *srv) createWearAmount(wearAmount dto.WearAmount) (*models.WearAmount, error) {

	wearAmountModel := &models.WearAmount{
		Description: wearAmount.Description,
	}

	err := s.DB.Create(&wearAmountModel).Error
	if err != nil {
		return nil, err
	}

	return wearAmountModel, nil
}

func (s *srv) listAllSkins() ([]*dto.SkinResponse, error) {
	var skins []*models.Skin
	err := s.DB.Find(&skins).Error
	if err != nil {
		return nil, err
	}

	var skinsResponse []*dto.SkinResponse
	for _, skin := range skins {

		skinResponse := &dto.SkinResponse{
			ID:        skin.ID,
			CreatedAt: skin.CreatedAt,
			Name:      skin.Name,
			ImageURL:  skin.ImageURL,
			Value:     skin.Value,
		}

		skinsResponse = append(skinsResponse, skinResponse)
	}

	return skinsResponse, nil
}

func (s *srv) createRarityType(rarityType dto.RarityType) (*models.RarityType, error) {

	rarityTypeModel := &models.RarityType{
		Description: rarityType.Description,
		Priority:    rarityType.Priority,
	}

	err := s.DB.Create(&rarityTypeModel).Error
	if err != nil {
		return nil, err
	}

	return rarityTypeModel, nil
}
