tests:
	go test -v ./domain/... -coverprofile coverage.out

coverage: tests
	go tool cover -html=coverage.out