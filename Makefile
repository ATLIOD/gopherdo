BINARY_NAME=gopherdo
BUILD_DIR=build

.PHONY: build
build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) api/gopherdo_api.go
	cp -r api/static $(BUILD_DIR)/

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: build-all
build-all: build-windows build-linux build-mac

.PHONY: build-windows
build-windows:
	mkdir -p $(BUILD_DIR)/windows
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/windows/$(BINARY_NAME).exe api/gopherdo_api.go
	cp -r api/static $(BUILD_DIR)/windows/

.PHONY: build-linux
build-linux:
	mkdir -p $(BUILD_DIR)/linux
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/linux/$(BINARY_NAME) api/gopherdo_api.go
	cp -r api/static $(BUILD_DIR)/linux/

.PHONY: build-mac
build-mac:
	mkdir -p $(BUILD_DIR)/mac
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/mac/$(BINARY_NAME) api/gopherdo_api.go
	cp -r api/static $(BUILD_DIR)/mac/

.PHONY: dist
dist: build-all
	cd $(BUILD_DIR) && \
	zip -r gopherdo-windows.zip windows && \
	zip -r gopherdo-linux.zip linux && \
	zip -r gopherdo-mac.zip mac
