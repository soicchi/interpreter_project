go_test:
	docker compose run --rm interpreter go test -v ./...

go_run:
	docker compose run --rm interpreter go run main.go
