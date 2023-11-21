// TODO: add commands for build and run in dev/produciton mode
// TODO: add commands for build protobuf files
// TODO: add commands for lint/format/test project

ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))


// TODO: add lint and format to github ci
lint:
	which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	golangci-lint run --config=$(ROOT)/.golangci.yml $(ROOT)/...