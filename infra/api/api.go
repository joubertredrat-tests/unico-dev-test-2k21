package api

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Run() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	fmt.Printf("listening http in %s", os.Getenv("API_LISTEN_PORT"))

	return nil
}
