package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/middleware"
	"github.com/poportss/jackportcs/internal/rest"
	"net/http"
)

func (s *srv) CreateUserTradeLinkHandler(c *gin.Context) {
	var userTradeLink dto.UserTradeLink

	if err := c.ShouldBindJSON(&userTradeLink); err != nil {
		rest.ResponseBadRequest(c, fmt.Errorf("Dados inválidos"))
		return
	}

	userID, err := middleware.ExtractUserIDFromContext(c)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	err = s.createUserTradeLink(userTradeLink.TradeLink, userID)
	if err != nil {
		rest.ResponseInternalServerError(c, err)
	}

	rest.ResponseDefaultSuccess(c)
	return
}

func (s *srv) CreateUserAddressHandler(c *gin.Context) {
	var userAddress dto.UserAddress

	if err := c.ShouldBindJSON(&userAddress); err != nil {
		rest.ResponseBadRequest(c, fmt.Errorf("Dados inválidos"))
		return
	}

	userID, err := middleware.ExtractUserIDFromContext(c)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	err = s.createUserAddress(userAddress, userID)
	if err != nil {
		rest.ResponseInternalServerError(c, err)
		return
	}

	rest.ResponseDefaultSuccess(c)
	return
}

func (s *srv) GetBalanceHandler(c *gin.Context) {
	userID, err := middleware.ExtractUserIDFromContext(c)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	balance, err := s.getBalance(userID)
	if err != nil {
		rest.ResponseInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
	return
}
