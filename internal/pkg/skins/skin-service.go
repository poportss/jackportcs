package skins

import (
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
)

func (s *srv) createSkin(skin dto.Skin) (*models.Skin, error) {

	skinModel := &models.Skin{
		Name:       skin.Name,
		ImageURL:   skin.ImageURL,
		DropChance: skin.DropChance,
		Value:      skin.Value,
	}

	err := s.DB.Create(&skinModel).Error
	if err != nil {
		return nil, err
	}

	return skinModel, nil
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
			ID:         skin.ID,
			CreatedAt:  skin.CreatedAt,
			Name:       skin.Name,
			ImageURL:   skin.ImageURL,
			DropChance: skin.DropChance,
			Value:      skin.Value,
		}

		skinsResponse = append(skinsResponse, skinResponse)
	}

	return skinsResponse, nil
}
