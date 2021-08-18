package log

import (
	"fmt"
	"log"
	"os"
)

func NewLogFile(filename string) (*log.Logger, error) {
	logFile, err := os.OpenFile(fmt.Sprintf("./var/%s", filename), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return log.New(logFile, "", log.LstdFlags), nil
}
