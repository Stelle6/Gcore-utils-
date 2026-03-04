SHELL := /bin/bash

UTILS := $(shell find . -mindepth 2 -maxdepth 2 -name go.mod -printf '%h\n' | sed 's#^\./##' | sort)
BIN_DIR := build/bin

.PHONY: all clean list $(UTILS)

all: $(UTILS)

$(UTILS):
	@mkdir -p $(BIN_DIR)
	@echo "Building $@ -> $(BIN_DIR)/$@"
	@cd $@ && go build -o ../$(BIN_DIR)/$@ .

list:
	@printf '%s\n' $(UTILS)

clean:
	rm -rf build
