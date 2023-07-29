default := all

.PHONY: all
all: build format

.PHONY: build
build:
	go build cmd/isitworththetime/isitworththetime.go

.PHONY: format
format:
	go fmt ./...

.PHONY: test
test: fuzz
	go test ./...

.PHONY: fuzz
fuzz:
	cd pkg/timebudget && go test -fuzz FuzzTimeSaved --fuzztime 15s

.PHONY: clean
clean:
	git clean -fxd