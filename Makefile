build:
	go build -o docker ./cmd/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o docker ./cmd/main.go