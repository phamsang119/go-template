package utils

import (
	"github.com/gin-gonic/gin"
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

func Reject(c *gin.Context, httpCode int, code int, message string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"status": gin.H{
			"code":    code,
			"message": message,
		},
		"result": gin.H{
			"data": data,
		},
	})
}
