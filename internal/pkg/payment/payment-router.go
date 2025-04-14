package payment

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/baseservice"
	"github.com/poportss/jackportcs/internal/migrations"
	"github.com/poportss/jackportcs/internal/pkg/payment/migration"
)

type srv struct {
	*baseservice.BaseService
	BraipApiBaseURL string
	BraipApiToken   string
}

func PaymentNewService(base *baseservice.BaseService) *srv {
	err := migrations.Migrate(base.DB, "payment", migration.Versions())
	if err != nil {
		panic("❌ Erro ao rodar as migrations: " + err.Error())
	}

	return &srv{BaseService: base}
}

func ConfigureRoutes(r *gin.Engine, base *baseservice.BaseService, jwtMiddleware *jwt.GinJWTMiddleware) {
	service := PaymentNewService(base)

	api := r.Group("/api")
	paymentRoutes := api.Group("/payment")
	paymentRoutes.Use(jwtMiddleware.MiddlewareFunc())
	{
		paymentRoutes.POST("/createPaymentOrder", service.CreatePaymentOrderHandler)
		paymentRoutes.POST("/createPaymentCustomer", service.CreatePaymentCustomerHandler)
		paymentRoutes.POST("/createCustomerCard", service.CreateCustomerCardHandler)
	}

	paymentPagarmeRoutes := api.Group("/payment/pagarme")
	paymentPagarmeRoutes.POST("webhook", service.PagarmeWebhookHandler)
}
