package api

import "github.com/gin-gonic/gin"

func NewGinEngine() *gin.Engine {
	r := gin.Default()

	r.GET("/", index)

	routes := r.Group("/api/v1")
	routes.POST("/packet", createPacket)
	routes.GET("/packet", packetList)

	return r
}
