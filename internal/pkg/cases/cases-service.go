package cases

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
	"gorm.io/gorm"
	"math"
	"math/rand"
	"time"
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
	err := s.DB.Where("id IN (?)", caseParam.SkinsID).Order("value desc").Find(&skins).Error
	if err != nil {
		return nil, err
	}

	if len(skins) == 0 {
		return nil, errors.New("nenhuma skin encontrada para a caixa")
	}

	// Calcular os pesos baseados no inverso do valor
	var weightTotal float64
	weights := make([]float64, len(skins))
	for i, skin := range skins {
		if skin.Value <= 0 {
			return nil, fmt.Errorf("valor inválido para a skin: %s", skin.ID)
		}
		// Quanto maior o valor, menor o peso
		weights[i] = 1.0 / float64(skin.Value)
		weightTotal += weights[i]
	}

	// Calcular as probabilidades
	var probabilities []map[string]interface{}
	var sum int
	for i, skin := range skins {
		probability := math.Round((weights[i] / weightTotal) * 100)
		sum += int(probability)
		probabilities = append(probabilities, map[string]interface{}{
			"skinID":      skin.ID,
			"probability": int(probability),
		})
	}

	// Ajustar a última probabilidade para somar 100
	if sum != 100 && len(probabilities) > 0 {
		delta := 100 - sum
		lastIndex := len(probabilities) - 1
		probabilities[lastIndex]["probability"] = probabilities[lastIndex]["probability"].(int) + delta
	}

	// Converter para JSON
	probabilitiesJSON, err := json.Marshal(probabilities)
	if err != nil {
		return nil, err
	}
	caseModel.Probabilities = probabilitiesJSON

	// Salvar no banco
	err = s.DB.Create(&caseModel).Error
	if err != nil {
		return nil, err
	}

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
func (s *srv) openCaseByID(caseID string, userID uuid.UUID) (*dto.SkinOpenedCase, error) {
	var skinOpenedCaseResponse *dto.SkinOpenedCase

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		var caseData models.Case
		if err := tx.Where("id = ?", caseID).First(&caseData).Error; err != nil {
			return fmt.Errorf("erro ao buscar o case: %w", err)
		}

		var user models.User
		if err := tx.Where("id = ?", userID).Preload("Wallet").First(&user).Error; err != nil {
			return fmt.Errorf("erro ao buscar usuário: %w", err)
		}

		if user.Wallet == nil || user.Wallet.Balance < caseData.Price {
			return errors.New("saldo insuficiente")
		}

		if err := json.Unmarshal(caseData.Probabilities, &caseData.CaseProbability); err != nil {
			return fmt.Errorf("erro ao processar probabilidades: %w", err)
		}

		if len(caseData.CaseProbability) == 0 {
			return errors.New("nenhuma probabilidade definida para o case")
		}

		rand.Seed(time.Now().UnixNano())
		random := rand.Int63n(101)
		var selectedSkin uuid.UUID
		var cumulativeProbability int64

		for _, prop := range caseData.CaseProbability {
			cumulativeProbability += prop.Probability
			if random < cumulativeProbability {
				selectedSkin = prop.SkinID
				break
			}
		}

		// Fallback
		if selectedSkin == uuid.Nil {
			fallback := caseData.CaseProbability[rand.Intn(len(caseData.CaseProbability))]
			selectedSkin = fallback.SkinID
		}

		user.Wallet.Balance -= caseData.Price
		if err := tx.Save(&user.Wallet).Error; err != nil {
			return fmt.Errorf("erro ao atualizar saldo: %w", err)
		}

		// Registrar transação do usuário
		userTransaction := models.UserTransaction{
			UserID:       userID,
			Type:         models.TransactionTypeCaseOpen,
			Description:  fmt.Sprintf("Abertura de caixa: %s", caseData.Name),
			Amount:       -caseData.Price,
			BalanceAfter: user.Wallet.Balance,
		}
		if err := tx.Create(&userTransaction).Error; err != nil {
			return fmt.Errorf("erro ao registrar transação: %w", err)
		}

		inventorySkin := models.InventorySkins{
			InventoryID: user.InventoryID,
			SkinID:      selectedSkin,
		}
		if err := tx.Create(&inventorySkin).Error; err != nil {
			return fmt.Errorf("erro ao adicionar skin ao inventário: %w", err)
		}

		var skin models.Skin
		if err := tx.Preload("WearAmount").Preload("RarityType").
			Where("id = ?", selectedSkin).First(&skin).Error; err != nil {
			return fmt.Errorf("erro ao buscar skin: %w", err)
		}

		skinOpenedCaseResponse = &dto.SkinOpenedCase{
			ID:              skin.ID,
			Name:            skin.Name,
			ImageURL:        skin.ImageURL,
			WearAmount:      skin.WearAmount.Description,
			RarityType:      skin.RarityType.Description,
			Value:           skin.Value,
			InventorySkinID: inventorySkin.ID,
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return skinOpenedCaseResponse, nil
}
