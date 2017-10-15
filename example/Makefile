GO  			?= go
BUILDFLAGS		:= -v -tags 'netgo'
TESTFLAGS		:= -cover
BINDIR			:= bin

DOCKER_COMPOSE	?= docker-compose

export GOOS        := $(subst Darwin,darwin,$(subst Linux,linux,$(subst FreeBSD,freebsd,$(shell uname))))
export CGO_ENABLED := 0
export GOGC        := 400

.PHONY: build
build:
	$(GO) build $(BUILDFLAGS)

.PHONY: clean
clean:
	@$(GO) clean
	@rm -rf $(BINDIR)/*

.PHONY: test
test: PKG ?= $(shell go list./... | grep -v /vendor/)
test:
	@$(GO) test $(TESTFLAGS) $(PKG)

.PHONY: docker-run
docker-run: export DOCKER_VOLUME := /go/src/github.com/heltonmarx/goami/example
docker-run:
	@$(DOCKER_COMPOSE) run --rm -w $(DOCKER_VOLUME) -e DOCKER_VOLUME=$(DOCKER_VOLUME) ami-test /bin/bash

