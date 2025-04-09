package user

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/migrations"
	"github.com/poportss/jackportcs/internal/pkg/user/migration"
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

func ConfigureRoutes(r *gin.Engine, base *baseservice.BaseService, jwtMiddleware *jwt.GinJWTMiddleware) {
	service := UserNewService(base)
	api := r.Group("/api")

	casesRoutes := api.Group("/user")
	casesRoutes.Use(jwtMiddleware.MiddlewareFunc())
	{
		casesRoutes.POST("/createUserTradeLink", service.CreateUserTradeLinkHandler)
	}

}
