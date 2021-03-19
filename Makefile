GOBUILD = go build
BINARY_NAME = countries

GOBASE = $(shell pwd)
CMD = $(GOBASE)/cmd
GOFILES = $(CMD)/countries/*.go
GOBIN = $(GOBASE)/bin

all: build

deb:

build:
	$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) $(GOFILES)

clean:
	rm $(GOBIN)/$(BINARY_NAME)