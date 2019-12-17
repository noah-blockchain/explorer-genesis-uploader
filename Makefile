APP ?= explorer-uploader
GOOS ?= linux
SRC = ./

all: test build

#Run this from CI
create_vendor:
	@rm -rf vendor
	@echo "--> Running go mod vendor"
	@go mod vendor

### Build ###################
build: clean
	GOOS=${GOOS} go build -o ./build/$(APP) -i ./cmd/explorer-uploader

install:
	GOOS=${GOOS} go install -i ./cmd/explorer-uploader

clean:
	@rm -f $(BINARY)

### Test ####################
test:
	@echo "--> Running tests"
	go test -v ${SRC}

fmt:
	@go fmt ./...

.PHONY: create_vendor build clean fmt test
