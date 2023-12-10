package handlers

import (
	"net/http"

	"github.com/GoTH/components"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "", components.Index())
}