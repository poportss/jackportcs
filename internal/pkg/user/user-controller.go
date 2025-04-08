package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/middleware"
	"github.com/poportss/jackportcs/internal/rest"
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
