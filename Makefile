GIT_REV := "$(shell git rev-parse --short HEAD )"
GO := go
GO_BUILD := $(GO) build
GO_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

SUFFIX := .go

.DEFAULT_GOAL := help
.PHONY: all handlers devdeps clean help

# Register dependency files to detect changes
DEPFILES :=
LIBS := \

DEPFILES += $(LIBS:%=%/*$(SUFFIX))

LAMBDA_HANDLER_DIR := handler
LAMBDA_HANDLERS := \
	register
DEPFILES += $(addprefix $(LAMBDA_HANDLER_DIR)/, $(LAMBDA_HANDLERS:%=%/*$(SUFFIX)))

DIST_DIR := _dist
TARGETS := $(LAMBDA_HANDLERS:%=$(DIST_DIR)/%)

$(DIST_DIR)/%: $(LAMBDA_HANDLER_DIR)/% $(DEPFILES) go.sum
	$(GO_ENV) $(GO_BUILD) -ldflags="-s -w -X main.Version=$(GIT_REV)" -o $@ ./$<

## Build all handlers
handlers: $(TARGETS)

## Install dependencies for development
devdeps:
	GO111MODULE=off go get \
	github.com/Songmu/make2help/cmd/make2help

## Clean up artifacts
clean:
	$(GO) clean
	rm -rf $(DIST_DIR)

## Show help
help:
	@make2help $(MAKEFILE_LIST)
