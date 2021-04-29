PROJECT_NAME=$(shell basename "$(PWD)")
BIN_DIR=bin

build-client: clean
	go build -o ${BIN_DIR}/$(PROJECT_NAME)-client-linux cmd/client/main.go
	GOOS=windows GOARCH=amd64 go build -o ${BIN_DIR}/$(PROJECT_NAME)-client-windows.exe cmd/client/main.go

build-server: clean
	go build -o ${BIN_DIR}/$(PROJECT_NAME)-client-linux cmd/server/main.go

build-all: build-clean build-server

clean:
	rm -rf bin/
