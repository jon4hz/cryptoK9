BINARY_NAME=cryptoK9

run:
	@go run ./cmd/cryptoK9 ||:

build:
	go build -o ${BINARY_NAME} -ldflags="-s" cmd/cryptoK9/main.go

docker-push:
	docker build -t jon4hz/cryptok9:latest .
	docker push jon4hz/cryptok9:latest

docker-build:
	docker build -t jon4hz/cryptok9:latest .

docker:
	docker-compose up -d	

test:
	go test ./...