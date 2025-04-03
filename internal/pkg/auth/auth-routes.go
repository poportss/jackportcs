package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/migrations"
	"github.com/poportss/jackportcs/internal/pkg/auth/migrations"
	"gorm.io/gorm"
)

type srv struct {
	*baseservice.BaseService
}

// NewMyService cria um novo serviço, herdando o base service
func AuthNewService(base *baseservice.BaseService) *srv {
	err := migrations.Migrate(base.DB, "auth", migration.Versions())
	if err != nil {
		panic("❌ Erro ao rodar as migrations: " + err.Error())
	}

	return &srv{BaseService: base}
}

// RegisterAuthRoutes adiciona as rotas de autenticação ao router principal
func ConfigureRoutes(r *gin.Engine, db *gorm.DB) {
	handler := srv{baseservice.NewBaseService(db)}

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", handler.LoginHandler)
	}
}
