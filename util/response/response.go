package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Info(c *gin.Context, httpStatus int, code int, message string, data interface{}) {
	c.JSON(httpStatus, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    SUCCESS,
		"message": "success",
		"data":    data,
	})
}

func Failed(c *gin.Context, err string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    FAILED,
		"message": "failed",
		"data":    err,
	})
}

func Error(c *gin.Context, err string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    Server_Error,
		"message": "server error",
		"data":    err,
	})
}
