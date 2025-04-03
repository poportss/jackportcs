package auth

import (
	"fmt"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
	"github.com/poportss/jackportcs/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *srv) login(login dto.Login) (string, error) {
	var user models.User
	if err := s.DB.Where("username = ?", login.Username).First(&user).Error; err != nil {
		return "", fmt.Errorf("Usuário não encontrado")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return "", fmt.Errorf("Senha incorreta")
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		return "", fmt.Errorf("Erro ao gerar token")
	}

	return token, nil
}
