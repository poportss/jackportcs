package skins

import (
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/dto"
	"net/http"
)

func (s *srv) CreateSkinHandler(c *gin.Context) {
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

func (s *srv) CreateWearAmountHandler(c *gin.Context) {
	var wearAmount dto.WearAmount

	if err := c.ShouldBindJSON(&wearAmount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	skinModel, err := s.createWearAmount(wearAmount)
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

func (s *srv) CreateRarityTypeHandler(c *gin.Context) {
	var rarityType dto.RarityType

	if err := c.ShouldBindJSON(&rarityType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	skinModel, err := s.createRarityType(rarityType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, skinModel)
}
