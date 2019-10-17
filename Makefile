GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=faasup

all: clean deps build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/spf13/cobra


