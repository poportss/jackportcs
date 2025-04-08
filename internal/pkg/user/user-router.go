package user

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/migrations"
	"github.com/poportss/jackportcs/internal/pkg/user/migration"
	"gorm.io/gorm"
)

type srv struct {
	*baseservice.BaseService
}

func UserNewService(base *baseservice.BaseService) *srv {
	err := migrations.Migrate(base.DB, "user", migration.Versions())
	if err != nil {
		panic("❌ Erro ao rodar as migrations: " + err.Error())
	}

	return &srv{BaseService: base}
}

func ConfigureRoutes(r *gin.Engine, db *gorm.DB, jwtMiddleware *jwt.GinJWTMiddleware) {
	handler := srv{baseservice.NewBaseService(db)}

	api := r.Group("/api")

	casesRoutes := api.Group("/user")
	casesRoutes.Use(jwtMiddleware.MiddlewareFunc())
	{
		casesRoutes.POST("/createUserTradeLink", handler.CreateUserTradeLinkHandler)
	}

}
