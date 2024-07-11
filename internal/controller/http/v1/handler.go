package v1

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	v1 := router.Group("/api/v1", h.userIdentity)

	client := v1.Group("/client")
	{
		client.POST("", h.addClient)
		client.PUT("/update", h.updateClient)
		client.DELETE("/delete", h.deleteClient)
	}

	status := v1.Group("/status")
	{
		status.PUT("/update", h.updateAlgorithmStatus)
	}

	return router
}
