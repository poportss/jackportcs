package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/migrations"
	"github.com/poportss/jackportcs/internal/pkg/auth/migration"
	"net/http"
)

type srv struct {
	*baseservice.BaseService
}

// NewMyService cria um novo serviço, herdando o base service
func AuthNewService(base *baseservice.BaseService) *srv {
	return &srv{BaseService: base}
}

// RegisterAuthRoutes adiciona as rotas de autenticação ao router principal
func ConfigureRoutes(r *gin.Engine, base *baseservice.BaseService, jwtMiddleware *jwt.GinJWTMiddleware) {
	err := migrations.Migrate(base.DB, "auth", migration.Versions())
	if err != nil {
		panic("❌ Erro ao rodar as migrations: " + err.Error())
	}
	service := AuthNewService(base)
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", jwtMiddleware.LoginHandler)
		authRoutes.GET("/refresh_token", jwtMiddleware.RefreshHandler)

		authRoutes.POST("/steam/login", service.SteamLoginHandler(jwtMiddleware))

		authRoutes.GET("me", service.Me)

		authRoutes.Use(jwtMiddleware.MiddlewareFunc())
		{
			authRoutes.GET("/validate", func(c *gin.Context) {
				claims := jwt.ExtractClaims(c)
				c.JSON(http.StatusOK, gin.H{"claims": claims})
			})
		}
	}

}
