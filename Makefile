BINARY_NAME = countries
GOBUILD = go build
GOTEST = go test

GOBASE = $(shell pwd)
CMD = $(GOBASE)/cmd
GOFILES = $(CMD)/countries/*.go
GOBIN = $(GOBASE)/bin

GOTESTFILES = $(GOBASE)/internal/country/*.go

all: build

build:
	$(GOBUILD) -o $(GOBIN)/$(BINARY_NAME) $(GOFILES)

test:
	$(GOTEST) $(GOTESTFILES)

clean:
	rm $(GOBIN)/$(BINARY_NAME)
