
VERSION := v0.0.1

GOBIN := /Users/loy/Plugins/go

publish: build cp2go
	@echo "success"

build:
	go build -o release/$(VERSION)/gdm ./main.go

cp2go:
	cp release/$(VERSION)/gdm  $(GOBIN)/bin/