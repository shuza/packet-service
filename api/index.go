package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, "Packer service is up and running...")
}
