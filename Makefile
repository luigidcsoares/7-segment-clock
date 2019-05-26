# Go related variables.
GOBASE = $(shell pwd)
GOBIN = $(GOBASE)/bin
GOFILES = $(wildcard $(GOBASE)/src/*.go)

# Name of the generated binary
PROJECTNAME = "clock"

# Make is verbose on linux. Make it silent.
MAKEFLAGS += --silent

all: go-build run

run:
	$(GOBIN)/$(PROJECTNAME)

go-build:
	go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-run:
	go run $(GOFILES)
