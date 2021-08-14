BINARY_NAME=mnemonicK9

run:
	@go run ./cmd/mnemonicK9 ||:

build:
	go build -o ${BINARY_NAME} -ldflags="-s" cmd/mnemonicK9/main.go

docker-push:
	docker build -t jon4hz/mnemonick9:latest .
	docker push jon4hz/mnemonick9:latest

docker:
	docker-compose up -d	

test:
	go test ./...