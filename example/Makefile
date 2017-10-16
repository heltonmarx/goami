GO  			?= go
BUILDFLAGS		:= -v -tags 'netgo'
TESTFLAGS		:= -cover
TESTPKG         := $(shell go list ./... | grep -v /vendor/)

export GOOS        := $(subst Darwin,darwin,$(subst Linux,linux,$(subst FreeBSD,freebsd,$(shell uname))))
export CGO_ENABLED := 0
export GOGC        := 400

.PHONY: build
build:
	$(GO) build $(BUILDFLAGS)

.PHONY: test
test:
	@$(GO) test $(TESTFLAGS) $(TESTPKG)

.PHONY: clean
clean:
	@$(GO) clean

