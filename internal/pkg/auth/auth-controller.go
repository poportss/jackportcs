package auth

import (
	"github.com/poportss/jackportcs/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *srv) LoginHandler(c *gin.Context) {
	var login dto.Login

	// Valida os dados enviados
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	token, err := s.login(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
