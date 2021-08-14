BINARY_NAME=mnemonicK5

run:
	@go run ./cmd/mnemonicK5 ||:

build:
	go build -o ${BINARY_NAME} -ldflags="-s" cmd/mnemonicK5/main.go

docker-push:
	docker build -t jon4hz/mnemonick5:latest .
	docker push jon4hz/mnemonick5:latest

docker:
	docker-compose up -d	

test:
	go test ./...