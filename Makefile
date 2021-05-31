# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.
.PHONY: gcodb

GOBIN = $(shell pwd)/build/bin
GO ?= latest

gcodb:
	build/env.sh go run build/ci.go install ./cmd/jinbao
	@echo "Done building."
	@echo "Run \"$(GOBIN)/jinbao\" to launch jinbao."

lint: ## Run linters.
	build/env.sh go run build/ci.go lint

clean:
	./build/clean_go_build_cache.sh
