// TODO: add commands for build and run in dev/produciton mode

ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

OS := $(shell uname -s)

lint:
	which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	golangci-lint run --config=$(ROOT)/.golangci.yml $(ROOT)/...

test:
	go test ./...

docker-test-up:
	docker compose -f $(ROOT)/deployment/test/docker-compose.yml up -d

docker-test-down:
	docker compose -f $(ROOT)/deployment/test/docker-compose.yml down

docker-local-up:
	sh -c "$(ROOT)/deployment/local/docker-compose.bash up -d"

logs:
	docker compose -f $(ROOT)/deployment/test/docker-compose.yml logs

format:
	@which gofumpt || (go install mvdan.cc/gofumpt@latest)
	@gofumpt -l -w $(ROOT)
	@which gci || (go install github.com/daixiang0/gci@latest)
	@gci write $(ROOT)
	@which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	@golangci-lint run --fix

protobuf:
	find contract/protobuf/ -name '*.proto' | xargs -I {} protoc --go-grpc_out=contract/go/ --go-grpc_opt=paths=source_relative --go_out=contract/go --go_opt=paths=source_relative --proto_path=contract/protobuf/ {}

swagger-gen:
	@which swag || (go install github.com/swaggo/swag/cmd/swag@latest)
	swag fmt
	swag init -g ../cmd/manager/main.go -d manager/ --parseDependency --output doc/swagger --instanceName manager
	swag init -g ../cmd/source/main.go -d source/ --parseDependency --output doc/swagger --instanceName source
	docker compose --env-file ./deployment/local/.env -f ./deployment/local/services/swagger.yml restart swagger

drop-manager-db:
	docker compose --env-file ./deployment/local/.env -f ./deployment/local/services/scylladb.yml exec scylladb cqlsh -e "DROP KEYSPACE manager;"