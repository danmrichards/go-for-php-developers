GOARCH=amd64
BINARY=hello
GO111MODULE=on

build: linux darwin windows

linux:
		GO111MODULE=${GO111MODULE} CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=linux go build -o ./bin/${BINARY}-linux-${GOARCH} .

darwin:
		GO111MODULE=${GO111MODULE} CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=darwin go build -o ./bin/${BINARY}-darwin-${GOARCH} .

windows:
		GO111MODULE=${GO111MODULE} CGO_ENABLED=0 GOARCH=${GOARCH} GOOS=windows go build -o ./bin/${BINARY}-windows-${GOARCH}.exe .

.PHONY: build