PROJECT_NAME=$(shell basename "$(PWD)")
BIN_DIR=bin
VERSION=$(shell cat VERSION)
LDFLAGS_CLIENT="-w -s -X client.main.serverURL=${SERVER_URL} -X client.main.token=${TOKEN}"

clean:
	rm -rf bin/

build-client: clean
	GOOS=linux GOARCH=amd64 go build -ldflags=${LDFLAGS_CLIENT} -o ${BIN_DIR}/${PROJECT_NAME}-client-linux-${VERSION} cmd/client/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags=${LDFLAGS_CLIENT} -o ${BIN_DIR}/${PROJECT_NAME}-client-windows-${VERSION}.exe cmd/client/main.go

build-server: clean
	GOOS=linux GOARCH=amd64 go build -o ${BIN_DIR}/$(PROJECT_NAME)-server-linux cmd/server/main.go

build-all: build-client build-server
