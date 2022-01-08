GOCMD=go
GOTEST=$(GOCMD) test
GOCOVER=$(GOCMD) tool cover
GOFMT=gofmt

.DEFAULT_GOAL := all

.PHONY: all
all: check-fmt test coverage

.PHONY: test
test:
	$(GOTEST) -v ./... -covermode=count -coverprofile=c.out

.PHONY: coverage
coverage:
	$(GOCOVER) -func=c.out

.PHONY: check-fmt
check-fmt:
	$(GOFMT) -d ${GOFILES}

.PHONY: fmt
fmt:
	$(GOFMT) -w ${GOFILES}