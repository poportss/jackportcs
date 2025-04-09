package auth

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/solovev/steam_go"
)

func SteamLoginHandler(jwtMiddleware *jwt.GinJWTMiddleware, base *baseservice.BaseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		openID := steam_go.NewOpenId(c.Request)

		switch openID.Mode() {
		case "":
			c.Redirect(http.StatusFound, openID.AuthUrl())
			return
		case "cancel":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Login cancelado pelo usuário"})
			return
		default:
			steamID, err := openID.ValidateAndGetId()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro validando Steam login"})
				return
			}

			service := AuthNewService(base)
			user, err := service.FindOrCreateUserBySteamID(steamID)
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
}
