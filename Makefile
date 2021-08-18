tests:
	go test -v ./domain/... -coverprofile coverage.out

coverage-html: tests ;
	go tool cover -html=coverage.out

coverage-console: tests ;
	go tool cover -func=coverage.out

up:
	docker-compose -f docker-compose.yml -p unico_dev_test up -d

status:
	docker-compose -f docker-compose.yml -p unico_dev_test ps

log:
	docker-compose -f docker-compose.yml -p unico_dev_test logs -f

down:
	docker-compose -f docker-compose.yml -p unico_dev_test down

coverage-docker:
	docker exec -it unico-dev-test-2k21_api.unico.dev_1 make coverage-console

logs-arquivo:
	docker exec -it unico-dev-test-2k21_api.unico.dev_1 tail -f ./var/app.log