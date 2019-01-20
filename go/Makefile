GOARCH=amd64
BINARY=hello

build: linux darwin windows

linux:
		CGO_ENABLED=0 GOOS=linux go build -o ./bin/${BINARY}-linux-${GOARCH} .

darwin:
		CGO_ENABLED=0 GOOS=darwin go build -o ./bin/${BINARY}-darwin-${GOARCH} .

windows:
		CGO_ENABLED=0 GOOS=windows go build -o ./bin/${BINARY}-windows-${GOARCH}.exe .

.PHONY: build