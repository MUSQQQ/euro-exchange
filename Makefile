run-local:
	@go run ./...

run-docker-local:
	@docker-compose -f docker-compose.yml up --build --force-recreate --remove-orphans

test:
	@go test -v ./...
