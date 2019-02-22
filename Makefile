GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=do-ddns-client

.PHONY: all test clean

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME)

build-all:
	$(GOBUILD) -o $(BINARY_NAME)
	GOOS=linux GOARCH=arm GOARM=5 $(GOBUILD) -o $(BINARY_NAME)_linux_arm5

test:
	$(GOTEST) ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)_linux_arm5