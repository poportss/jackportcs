package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/middleware"
	"github.com/poportss/jackportcs/internal/pkg/auth"
	"github.com/poportss/jackportcs/internal/pkg/cases"
	"github.com/poportss/jackportcs/internal/pkg/payment"
	"github.com/poportss/jackportcs/internal/pkg/skins"
	"github.com/poportss/jackportcs/internal/pkg/user"
	"log"
	"time"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(r *gin.Engine, baseService *baseservice.BaseService, braipApiBaseURL, braipApiToken string) {
	// Configura o middleware JWT
	jwtMiddleware, err := middleware.SetupJWTMiddleware([]byte(""), time.Hour, time.Hour*24, baseService)
	if err != nil {
		log.Fatalf("Erro ao configurar JWT: %v", err)
	}

	// Registrar rotas
	auth.ConfigureRoutes(r, baseService, jwtMiddleware)
	cases.ConfigureRoutes(r, baseService, jwtMiddleware)
	skins.ConfigureRoutes(r, baseService, jwtMiddleware)
	user.ConfigureRoutes(r, baseService, jwtMiddleware)
	payment.ConfigureRoutes(r, baseService, jwtMiddleware, braipApiBaseURL, braipApiToken)
}
