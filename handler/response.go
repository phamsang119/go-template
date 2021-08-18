package handler

import (
	"game-api/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Resolve(c *gin.Context, httpCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"status": gin.H{
			"code":    httpCode,
			"message": "success",
		},
		"result": gin.H{
			"data": data,
		},
	})
}

func RejectWithError(c *gin.Context, err entity.Error, data interface{}) {
	httpCode := http.StatusOK
	if err.Code < http.StatusNetworkAuthenticationRequired {
		httpCode = err.Code
	} else {
		httpCode = http.StatusBadRequest
	}
	c.JSON(httpCode, gin.H{
		"status": gin.H{
			"code":    err.Code,
			"message": err.Message,
		},
		"result": gin.H{
			"data": data,
		},
	})
}

func Reject(c *gin.Context, httpCode int, message string, data interface{}) {

	c.JSON(httpCode, gin.H{
		"status": gin.H{
			"code":    httpCode,
			"message": message,
		},
		"result": gin.H{
			"data": data,
		},
	})
}
