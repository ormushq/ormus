// TODO: add commands for build and run in dev/produciton mode
// TODO: add commands for build protobuf files

ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

lint:
	which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	golangci-lint run --config=$(ROOT)/.golangci.yml $(ROOT)/...

test:
	go test ./...

up:
	docker-compose -f ./deployment/test/docker-compose.yml up -d

down:
	docker-compose -f ./deployment/test/docker-compose.yml down

logs:
	docker-compose logs

format:
	@which gofumpt || (go install mvdan.cc/gofumpt@latest)
	@gofumpt -l -w $(ROOT)
	@which gci || (go install github.com/daixiang0/gci@latest)
	@gci write $(ROOT)
	@which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	@golangci-lint run --fix
