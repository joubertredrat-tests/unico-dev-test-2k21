package worker

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Run(filename string) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	fmt.Printf("running worker from %s", filename)
	return nil
}
