package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	auth := router.Group("/auth") 
	{
		auth.POST("/sign-in", h.signIn)
		auth.GET("/sign-out", h.signOut)
	}

	return  router
}