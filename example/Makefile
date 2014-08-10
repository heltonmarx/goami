GOCC=go

OS=linux
ARCH=amd64

ALL_PROGS := events login sippeers sms

all: $(ALL_PROGS)

events: events.go
	@echo "[CC] events.go"
	@$(GOCC) build -o events $^

login: login.go
	@echo "[CC] login.go"
	@$(GOCC) build -o login $^

sippeers: sippeers.go
	@echo "[CC] sippeers.go"
	@$(GOCC) build -o sippeers $^

sms: sms.go
	@echo "[CC] sms.go"
	@$(GOCC) build -o sms $^

format:
	@echo "[FORMAT] *.go files"
	@find . -name "*.go" -exec go fmt {} \;

clean:
	@echo "[CLEAN]"
	@rm -rf $(ALL_PROGS)
