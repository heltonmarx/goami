GO	?= go

ALL_PROGS := events login sippeers sms

all: $(ALL_PROGS)

events: events.go
	@echo "[GO] events.go"
	@$(GO) build -o events $^

login: login.go
	@echo "[GO] login.go"
	@$(GO) build -o login $^

sippeers: sippeers.go
	@echo "[GO] sippeers.go"
	@$(GO) build -o sippeers $^

sms: sms.go
	@echo "[GO] sms.go"
	@$(GO) build -o sms $^

clean:
	@echo "[CLEAN]"
	@rm -rf $(ALL_PROGS)
