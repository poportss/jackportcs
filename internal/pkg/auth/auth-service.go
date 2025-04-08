package auth

import (
	"fmt"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
	"gorm.io/gorm"
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
