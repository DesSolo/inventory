build:
	go build -o inventory-linux main.go
	GOOS=windows GOARCH=amd64 go build -o inventory-windows.exe main.go
