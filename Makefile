run:
	go run main.go serve

swag:
	swag init -g cmd/app/main.go

dry-build:
	go build -o app main.go
	rm app