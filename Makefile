GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=web-dashboard-api

all: build

build:
	$(GOGET) -u -v
	$(GOBUILD) -o $(BINARY_NAME) -v