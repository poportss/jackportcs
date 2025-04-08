package skins

import (
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/dto"
	"net/http"
)

func (s *srv) CreateCaseHandler(c *gin.Context) {
	var skin dto.Skin

	if err := c.ShouldBindJSON(&skin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	skinModel, err := s.createSkin(skin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, skinModel)
}

func (s *srv) ListAllSkinsHandler(c *gin.Context) {
	skins, err := s.listAllSkins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, skins)
}
