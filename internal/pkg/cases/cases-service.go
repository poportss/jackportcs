package cases

import (
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
)

func (s *srv) createCase(caseParam dto.Case) (*models.Case, error) {

	caseModel := &models.Case{
		Name:    caseParam.Name,
		Price:   caseParam.Price,
		SkinsID: caseParam.SkinsID,
		Active:  true,
	}

	err := s.DB.Create(&caseModel).Error
	if err != nil {
		return nil, err
	}

	return caseModel, nil
}

func (s *srv) listAllCases() ([]*dto.CaseResponse, error) {
	var cases []*models.Case
	err := s.DB.Preload("Skins").Find(&cases).Error
	if err != nil {
		return nil, err
	}

	var casesResponse []*dto.CaseResponse
	for _, caseUnit := range cases {

		caseResponse := &dto.CaseResponse{
			ID:        caseUnit.ID,
			CreatedAt: caseUnit.CreatedAt,
			Name:      caseUnit.Name,
			Price:     caseUnit.Price,
		}

		if caseUnit.Skins != nil && len(caseUnit.Skins) > 0 {
			caseResponse.Skins = caseUnit.Skins
		}

		casesResponse = append(casesResponse, caseResponse)
	}

	return casesResponse, nil
}
