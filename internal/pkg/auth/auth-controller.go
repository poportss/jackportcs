package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/dto"
	"net/http"
)

func (s *srv) SteamLoginHandler(jwtMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(c *gin.Context) {
		var steamIDRequest dto.SteamIDRequest

		if err := c.ShouldBindJSON(&steamIDRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		user, err := s.FindOrCreateUserBySteamID(steamIDRequest.SteamID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar/criar usuário"})
			return
		}

		c.Set("provider", "steam")
		c.Set("steamID", user.SteamID)
		c.Set(jwtMiddleware.IdentityKey, user)
		jwtMiddleware.LoginHandler(c)
	}
}

func (s *srv) Me(c *gin.Context) {
	steamID := c.Query("steamID")
	if steamID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "steamID não fornecido"})
		return
	}

	user, err := s.FindUserBySteamID(steamID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar/criar usuário"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
