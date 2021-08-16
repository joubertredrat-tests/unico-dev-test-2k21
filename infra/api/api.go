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
	router.GET("/api/health", controller.handleHealth)

	fmt.Printf("Listening http API in http://0.0.0.0:%s", os.Getenv("API_LISTEN_PORT"))
	router.Run(fmt.Sprintf(":%s", os.Getenv("API_LISTEN_PORT")))

	return nil
}
