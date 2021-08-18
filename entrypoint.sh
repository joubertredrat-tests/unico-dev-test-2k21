#!/bin/sh

go run main.go import --filename ./DEINFO_AB_FEIRASLIVRES_2014.csv
go run main.go api
tail -f ./var/$APP_LOG_FILENAME.log