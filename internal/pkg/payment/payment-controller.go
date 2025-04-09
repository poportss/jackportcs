package payment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/middleware"
	"github.com/poportss/jackportcs/internal/rest"
	"net/http"
)

func (s *srv) CreatePaymentOrderHandler(c *gin.Context) {
	var createOrderRequest *dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&createOrderRequest); err != nil {
		rest.ResponseBadRequest(c, fmt.Errorf("Dados inválidos"))
		return
	}

	userID, err := middleware.ExtractUserIDFromContext(c)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	paymentOrderResponse, err := s.createPaymentOrder(createOrderRequest, userID)
	if err != nil {
		rest.ResponseInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"paymentOrderResponse": paymentOrderResponse})
}

func (s *srv) CreatePaymentCustomerHandler(c *gin.Context) {
	var createOrderRequest *dto.PagarmeCreateCustomerRequest

	if err := c.ShouldBindJSON(&createOrderRequest); err != nil {
		rest.ResponseBadRequest(c, fmt.Errorf("Dados inválidos"))
		return
	}

	userID, err := middleware.ExtractUserIDFromContext(c)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	customer, err := s.createPaymentCustomer(createOrderRequest, userID)
	if err != nil {
		rest.ResponseInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"customer": customer})
}
