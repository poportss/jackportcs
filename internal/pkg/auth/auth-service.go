package auth

import (
	"fmt"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
	"github.com/poportss/jackportcs/internal/utils"
	"gorm.io/gorm"
	"os"
)

func (s *srv) AuthenticateUser(login dto.Login) (*models.User, error) {
	var user *models.User
	if err := s.DB.Where("name = ?", login.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, err
	}

	//if !utils.CheckPasswordHash(login.Password, user.Password) {
	//	return nil, fmt.Errorf("senha inválida")
	//}

	return user, nil
}

func (s *srv) FindOrCreateUserBySteamID(steamID string) (*models.User, error) {
	var user models.User
	if err := s.BaseService.DB.Where("steam_id = ?", steamID).First(&user).Error; err == nil {
		return &user, nil
	}

	// Fetch dados do perfil Steam
	name, avatar, err := utils.FetchSteamProfile(os.Getenv("STEAM_WEB_API_KEY"), steamID)
	if err != nil {
		name = "Jogador Steam"
	}

	// 💰 Cria carteira
	wallet := models.Wallet{}
	if err := s.BaseService.DB.Create(&wallet).Error; err != nil {
		return nil, err
	}

	// 💰 Cria inventário
	inventory := models.Inventory{}
	if err := s.BaseService.DB.Create(&inventory).Error; err != nil {
		return nil, err
	}

	// 👤 Cria usuário com wallet associada
	user = models.User{
		SteamID:   steamID,
		Name:      name,
		AvatarUrl: avatar,
		WalletID:  wallet.ID,
		Inventory: inventory,
	}

	if err := s.BaseService.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
