.PHONY: build run compile build-linux build-windows build-darwin test

VERSION ?= 0.2.0
NAME ?= "checkup"

test:
	go test ./tests -v

build:
	go build -o release/$(NAME) main.go

run:
	go run main.go

build-darwin:
	echo "Compiling for Darwin"
	GOOS=darwin GOARCH=amd64 go build -o release/$(NAME)-v${VERSION}-darwin-amd64 main.go
	# GOOS=darwin GOARCH=arm go build -v -o release/$(NAME)-v${VERSION}-darwin-arm main.go
	# GOOS=darwin GOARCH=arm64 go build -o release/$(NAME)-v${VERSION}-darwin-arm64 main.go

build-windows:
	echo "Compiling for Windows"
	GOOS=windows GOARCH=amd64 go build -o release/$(NAME)-v${VERSION}-windows-amd64 main.go
	GOOS=windows GOARCH=386 go build -o release/$(NAME)-v${VERSION}-windows-386 main.go

build-linux:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o release/$(NAME)-v${VERSION}-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o release/$(NAME)-v${VERSION}-linux-arm64 main.go
	GOOS=linux GOARCH=386 go build -o release/$(NAME)-v${VERSION}-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o release/$(NAME)-v${VERSION}-linux-amd64 main.go

compile: build-darwin build-windows build-linux