// TODO: add commands for build and run in dev/produciton mode
// TODO: add commands for build protobuf files

ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
LINT_BIN = $(GOPATH)/bin/golangci-lint

lint:
	which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	$(LINT_BIN) run --config=$(ROOT)/.golangci.yml $(ROOT)/...

test:
	go test -v ./...

format:
	@which gofumpt || (go install mvdan.cc/gofumpt@latest)
	@gofumpt -l -w $(ROOT)
	@which gci || (go install github.com/daixiang0/gci@latest)
	@gci write $(ROOT)
	@which $(LINT_BIN) || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	@$(LINT_BIN) run --fix