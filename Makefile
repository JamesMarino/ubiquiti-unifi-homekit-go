build:
	GOOS=linux GOARCH=mips64 go build -o homekit main.go
	chmod +x homekit
