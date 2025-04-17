package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
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

func GetMethod(url string) ([]byte, error) {
	var err error

	httpReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 30 * time.Second}

	res, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("error in response: %s", resBody)
	}

	return resBody, nil
}
