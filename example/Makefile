GO	?= go

ALL_PROGS := events login sippeers sms

all: $(ALL_PROGS)

events: events.go
	@echo "[CC] events.go"
	@$(GO) build -o events $^

login: login.go
	@echo "[CC] login.go"
	@$(GO) build -o login $^

sippeers: sippeers.go
	@echo "[CC] sippeers.go"
	@$(GO) build -o sippeers $^

sms: sms.go
	@echo "[CC] sms.go"
	@$(GO) build -o sms $^

format:
	@echo "[FORMAT] *.go files"
	@find . -name "*.go" -exec go fmt {} \;

clean:
	@echo "[CLEAN]"
	@rm -rf $(ALL_PROGS)
