package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseInternalServerError responds with a StatusInternalServerError, status FAILURE ad errorMessage
func ResponseInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"status": "FAILURE", "errorMessage": err.Error()})
}

// ResponseBadRequest responds with a StatusInternalServerError, status FAILURE ad errorMessage
func ResponseBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"status": "FAILURE", "errorMessage": err.Error()})
}

// ResponseDefaultSuccess responds with a Status200ok, status SUCCESS ad errorMessage empty
func ResponseDefaultSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "SUCCESS"})
}

func ResponseUnauthorize(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"status": "FAILURE", "errorMessage": "Unauthorized"})
	c.Abort()
}

func ResponseForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"status": "FAILURE", "errorMessage": "Forbidden"})
	c.Abort()
}

func ResponseNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"status": "FAILURE", "errorMessage": "Not Found"})
	c.Abort()
}
