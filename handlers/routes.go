package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	client := router.Group("/client")
	clientV1 := client.Group("/v1")
	clientV1.GET("/ports", h.ports)
	clientV1.POST("/ports", h.uploadPorts)

	return router
}
