package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/pkg/auth"
	"gorm.io/gorm"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	// Registrar rotas
	auth.ConfigureRoutes(r, db)
}
