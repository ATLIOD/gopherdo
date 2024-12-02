
.PHONY: build clean test run

BINARY_NAME=gopherdo
BUILD_DIR=build

VERSION ?= $(shell git describe --tags --always --dirty)
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS = -ldflags "-X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME}"


build:
	@echo "Building..."
	@mkdir -p ${BUILD_DIR}/{linux,mac,windows}/static
	# Linux
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/linux/${BINARY_NAME} ./cmd/server
	cp -r api/static/* ${BUILD_DIR}/linux/static/
	cd ${BUILD_DIR} && zip -r gopherdo-linux.zip linux
	# Mac
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/mac/${BINARY_NAME} ./cmd/server
	cp -r api/static/* ${BUILD_DIR}/mac/static/
	cd ${BUILD_DIR} && zip -r gopherdo-mac.zip mac
	# Windows
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o ${BUILD_DIR}/windows/${BINARY_NAME}.exe ./cmd/server
	cp -r api/static/* ${BUILD_DIR}/windows/static/
	cd ${BUILD_DIR} && zip -r gopherdo-windows.zip windows

clean:
	@echo "Cleaning..."
	@rm -rf ${BUILD_DIR}/*
	@rm -f data/*.db

run:
	go run ./cmd/server/main.go

test:
	go test -v ./...
