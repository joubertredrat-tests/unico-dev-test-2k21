package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	controller := NewController()
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/health", controller.handleHealth)
	api.GET("/open-markets", controller.handleListOpenMarket)
	api.POST("/open-markets", controller.handleCreateOpenMarket)
	api.GET("/open-markets/:id", controller.handleGetOpenMarket)
	api.PATCH("/open-markets/:id", controller.handleUpdateOpenMarket)
	api.DELETE("/open-markets/:id", controller.handleDeleteOpenMarket)

	fmt.Printf("Listening http API in http://0.0.0.0:%s", os.Getenv("API_LISTEN_PORT"))
	router.Run(fmt.Sprintf(":%s", os.Getenv("API_LISTEN_PORT")))

	return nil
}
