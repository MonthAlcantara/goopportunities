package router

import (
	"github.com/MonthAlcantara/goopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Definindo uma rota GET /ping
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.GetOpeningHandler)
	}
	{
		v1.POST("/opening", handler.CreateOpeningHandler)
	}
	{
		v1.DELETE("/opening", handler.DeleteOpeningHandler)

	}
	{
		v1.PUT("/opening", handler.UpdateOpeningHandler)
	}
	{
		v1.GET("/openings", handler.GetAllOpeningHandler)
	}
}
