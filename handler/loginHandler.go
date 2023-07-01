package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
