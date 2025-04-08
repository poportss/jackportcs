package skins

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/migrations"
	"github.com/poportss/jackportcs/internal/pkg/skins/migration"
	"gorm.io/gorm"
)

type srv struct {
	*baseservice.BaseService
}

func SkinNewService(base *baseservice.BaseService) *srv {
	err := migrations.Migrate(base.DB, "skin", migration.Versions())
	if err != nil {
		panic("❌ Erro ao rodar as migrations: " + err.Error())
	}

	return &srv{BaseService: base}
}

func ConfigureRoutes(r *gin.Engine, db *gorm.DB, jwtMiddleware *jwt.GinJWTMiddleware) {
	handler := srv{baseservice.NewBaseService(db)}

	api := r.Group("/api")

	skinRoutes := api.Group("/skin")
	skinRoutes.Use(jwtMiddleware.MiddlewareFunc())
	{
		skinRoutes.POST("/createSkin", handler.CreateCaseHandler)
		skinRoutes.GET("/listAllSkins", handler.ListAllSkinsHandler)
	}
}
