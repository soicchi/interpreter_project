go_test:
	docker compose run --rm interpreter go test -v -cover ./...

go_run:
	docker compose run --rm interpreter go run main.go
