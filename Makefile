PROJECT_NAME=$(shell basename "$(PWD)")

_mk_build_dir:
	@mkdir bin 

build: _mk_build_dir
	@echo "  >  Build linux"
	@go build -o bin/$(PROJECT_NAME)-linux main.go
	@echo "  >  Build windows"
	@GOOS=windows GOARCH=amd64 go build -o bin/$(PROJECT_NAME)-windows.exe main.go

clean:
	@rm -rf bin/
