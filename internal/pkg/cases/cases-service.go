package cases

import (
	"encoding/json"
	"errors"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
	"math"
)

func (s *srv) createCase(caseParam dto.Case) (*models.Case, error) {
	// Criar o modelo da caixa
	caseModel := &models.Case{
		Name:     caseParam.Name,
		Price:    caseParam.Price,
		Active:   true,
		ImageURL: caseParam.ImageURL,
	}

	// Buscar os valores das skins associadas à caixa
	var skins []models.Skin
	err := s.DB.Where("id IN (?)", caseParam.SkinsID).Order("value").Find(&skins).Error
	if err != nil {
		return nil, err
	}

	// Calcular o valor total das skins
	var totalValue int64
	for _, skin := range skins {
		totalValue += skin.Value
	}

	if totalValue == 0 {
		return nil, errors.New("valor total das skins é zero")
	}

	// Calcular as probabilidades das skins com base no preço
	var probabilities []map[string]interface{}
	var sum int
	for _, skin := range skins {
		// Calcular a probabilidade proporcional ao preço da skin
		probability := math.Round(float64(skin.Value) / float64(totalValue) * 100)
		sum += int(probability) // Acumular a soma das probabilidades

		probabilities = append(probabilities, map[string]interface{}{
			"skinID":      skin.ID,
			"probability": int(probability), // Armazenando a probabilidade como inteiro
		})
	}

	// Ajustar a última probabilidade para garantir que a soma seja 100
	if sum != 100 {
		delta := 100 - sum
		// Ajustar a última probabilidade (ou qualquer lógica desejada)
		lastIndex := len(probabilities) - 1
		probabilities[lastIndex]["probability"] = int(probabilities[lastIndex]["probability"].(float64)) + delta
	}

	// Converter o slice de probabilidades para JSON
	probabilitiesJSON, err := json.Marshal(probabilities)
	if err != nil {
		return nil, err
	}

	caseModel.Probabilities = probabilitiesJSON

	err = s.DB.Create(&caseModel).Error
	if err != nil {
		return nil, err
	}

	// Retornar o modelo da caixa criada
	return caseModel, nil
}

func (s *srv) listAllCases() ([]*dto.CaseResponse, error) {
	var cases []*models.Case
	err := s.DB.Where("active = ?", true).Find(&cases).Error
	if err != nil {
		return nil, err
	}

	var casesResponse []*dto.CaseResponse
	for _, caseUnit := range cases {

		caseResponse := &dto.CaseResponse{
			ID:        caseUnit.ID,
			CreatedAt: caseUnit.CreatedAt,
			ImageURL:  caseUnit.ImageURL,
			Name:      caseUnit.Name,
			Price:     caseUnit.Price,
		}

		casesResponse = append(casesResponse, caseResponse)
	}

	return casesResponse, nil
}

func (s *srv) getCaseByID(caseID string) (*dto.CaseDetailsResponse, error) {
	var caseUnit models.Case

	err := s.DB.Where("id = ?", caseID).First(&caseUnit).Error
	if err != nil {
		return nil, err
	}

	var probabilities []models.CaseProbability
	if err := json.Unmarshal(caseUnit.Probabilities, &probabilities); err != nil {
		return nil, err
	}

	response := &dto.CaseDetailsResponse{
		ID:        caseUnit.ID,
		CreatedAt: caseUnit.CreatedAt,
		Name:      caseUnit.Name,
		ImageURL:  caseUnit.ImageURL,
		Price:     caseUnit.Price,
	}

	for _, prob := range probabilities {
		var skin models.Skin
		if err := s.DB.Where("id = ?", prob.SkinID).Preload("RarityType").First(&skin).Error; err != nil {
			return nil, err
		}

		response.CaseProbabilities = append(response.CaseProbabilities, dto.CaseProbabilities{
			Skin: dto.SkinMetadataResponse{
				ID:        skin.ID,
				CreatedAt: skin.CreatedAt,
				ImageURL:  skin.ImageURL,
				Name:      skin.Name,
				Value:     skin.Value,
				Rarity:    skin.RarityType.Description,
			},
			Probability: prob.Probability,
		})
	}

	return response, nil
}
