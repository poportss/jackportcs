package cases

import (
	"github.com/gin-gonic/gin"
	"github.com/poportss/jackportcs/internal/dto"
	"github.com/poportss/jackportcs/internal/middleware"
	"github.com/poportss/jackportcs/internal/rest"
	"net/http"
)

func (s *srv) CreateCaseHandler(c *gin.Context) {
	var caseParam dto.Case

	if err := c.ShouldBindJSON(&caseParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	caseModel, err := s.createCase(caseParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, caseModel)
}

func (s *srv) ListAllCasesHandler(c *gin.Context) {
	allCases, err := s.listAllCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, allCases)
}

func (s *srv) GetCaseByIDHandler(c *gin.Context) {
	caseID := c.Param("ID")

	caseUnit, err := s.getCaseByID(caseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, caseUnit)
}

func (s *srv) OpenCaseByIDHandler(c *gin.Context) {
	caseID := c.Param("ID")

	userID, err := middleware.ExtractUserIDFromContext(c)
	if err != nil {
		rest.ResponseBadRequest(c, err)
		return
	}

	caseUnit, err := s.openCaseByID(caseID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, caseUnit)
}
