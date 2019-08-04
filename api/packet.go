package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"packet-service/model"
)

func createPacket(c *gin.Context) {
	var packet model.Packet
	if err := c.BindJSON(&packet); err != nil {
		log.Warnf("/create packet can't parse request body Error :  %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "can't parse request body",
			"data":    err.Error(),
		})

		return
	}

}
