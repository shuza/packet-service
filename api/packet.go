package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"packet-service/db"
	"packet-service/model"
	"packet-service/service"
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

	boxService := service.NewBoxService(service.NewHttpClient(), os.Getenv("BOX_SERVICE_ADDRESS"))
	box, err := boxService.FindSuitableBox(len(packet.Items), packet.Weight)
	if err != nil {
		log.Warnf("/create packet can't find suitable box Error :  %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Can't find suitable box",
			"data":    err.Error(),
		})
		return
	}

	packet.BoxId = box.Id
	if err := db.Client.Create(packet); err != nil {
		log.Warnf("/create packet can't save packet Error :  %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Can't save in Database",
			"data":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successful",
		"data":    packet,
	})
}
