package router

import (
	//lib to create the routes
	"github.com/gin-gonic/gin"
)

func NewConection() {
	//Initialize router with Gin's default config
	router := gin.Default()
	// Initialize routes
	initializeRoutes(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 ("localhost:8080")
}
