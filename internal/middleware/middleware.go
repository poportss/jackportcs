package middleware

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/models"
	"github.com/poportss/jackportcs/internal/pkg/auth"
	"time"
)

func SetupJWTMiddleware(jwtKey []byte, timeout, maxRefresh time.Duration, baseService *baseservice.BaseService) (*jwt.GinJWTMiddleware, error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "jackportcs",
		Key:         jwtKey,
		Timeout:     timeout,
		MaxRefresh:  maxRefresh,
		IdentityKey: "id",

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"id":   user.ID,
					"name": user.Name,
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Base: models.Base{
					ID: uuid.FromStringOrNil(claims["id"].(string)),
				},
			}
		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login dto.Login
			if err := c.ShouldBind(&login); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			provider, _ := c.Get("provider")

			var err error

			authService := auth.AuthNewService(baseService)

			var authenticateUser *models.User
			switch provider {
			case "steam":
				steamID, _ := c.Get("steamID")
				authenticateUser, err = authService.FindOrCreateUserBySteamID(steamID.(string))
				if err != nil {
					return nil, err
				}
			default:
				authenticateUser, err = authService.AuthenticateUser(login)
				if err != nil {
					return nil, err
				}
			}

			return authenticateUser, err
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"error": message})
		},

		TokenLookup: "header: Authorization, query: token",
	})
}

func ExtractUserIDFromContext(c *gin.Context) (uuid.UUID, error) {
	claims := jwt.ExtractClaims(c)

	userIDStr, ok := claims["id"].(string)
	if !ok || userIDStr == "" {
		return uuid.Nil, fmt.Errorf("user id não encontrado ou inválido")
	}

	userID := uuid.FromStringOrNil(userIDStr)
	if userID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("user id inválido")
	}

	return userID, nil
}
