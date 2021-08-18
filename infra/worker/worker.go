package worker

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	domainService "github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/service"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/infra/domain/fair/repository"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/infra/log"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/infra/mysql"
)

func Run(filename string) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	log, err := log.NewLogFile(os.Getenv("APP_LOG_FILENAME"))
	if err != nil {
		return err
	}

	if !fileExists(filename) {
		return errors.New(fmt.Sprintf("File %s does not exist", filename))
	}

	csvFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
		return err
	}

	db, err := mysql.NewMysqlConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	openMarketRepositoryMysql := repository.NewOpenMarketRepositoryMysql(db, log)
	openMarketService := domainService.NewOpenMarketService(openMarketRepositoryMysql)

	fmt.Printf("Start importing open markets from %s\n", filename)

	countSuccess := 0
	countError := 0
	csvLines = csvLines[1:]
	for _, line := range csvLines {
		openMarket := createOpenMarketFromCSVLine(line)
		_, err := openMarketService.Create(openMarket)
		if err != nil {
			countError++
			continue
		}

		countSuccess++
	}

	fmt.Printf("Finish import open markets from %s: success %d, errors %d\n", filename, countSuccess, countError)
	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
