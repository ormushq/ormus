// TODO: add commands for build and run in dev/produciton mode

ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

lint:
	which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.0)
	golangci-lint run --config=$(ROOT)/.golangci.yml $(ROOT)/...

test:
	go test ./...

docker-test-up:
	docker compose -f $(ROOT)/deployment/test/docker-compose.yml up -d

docker-test-down:
	docker compose -f $(ROOT)/deployment/test/docker-compose.yml down

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
ifneq (,$(findstring NT,$(OS)))
	protoc --go-grpc_out=contract/go/ --go-grpc_opt=paths=source_relative  --go_out=contract/go --go_opt=paths=source_relative --proto_path=./contract/protobuf/ contract/protobuf/event/event.proto
	protoc --go-grpc_out=contract/go/ --go-grpc_opt=paths=source_relative  --go_out=contract/go --go_opt=paths=source_relative --proto_path=./contract/protobuf/ contract/protobuf/task/task.proto
	protoc --go-grpc_out=contract/go/ --go-grpc_opt=paths=source_relative  --go_out=contract/go --go_opt=paths=source_relative --proto_path=./contract/protobuf/ contract/protobuf/internalevent/internalevent.proto
	protoc --go-grpc_out=contract/go/ --go-grpc_opt=paths=source_relative  --go_out=contract/go --go_opt=paths=source_relative --proto_path=./contract/protobuf/ contract/protobuf/project/project.proto
	protoc --go-grpc_out=contract/go/ --go-grpc_opt=paths=source_relative  --go_out=contract/go --go_opt=paths=source_relative --proto_path=./contract/protobuf/ contract/protobuf/source/source.proto
	protoc --go-grpc_out=contract/go/ --go-grpc_opt=paths=source_relative  --go_out=contract/go --go_opt=paths=source_relative --proto_path=./contract/protobuf/ contract/protobuf/user/user.proto
	protoc --go-grpc_out=contract/go/ --go-grpc_opt=paths=source_relative  --go_out=contract/go --go_opt=paths=source_relative --proto_path=./contract/protobuf/ contract/protobuf/user/user_info.proto
else
	find contract/protobuf/ -name '*.proto' | xargs -I {} protoc --go-grpc_out=contract/go/ --go-grpc_opt=paths=source_relative --go_out=contract/go --go_opt=paths=source_relative --proto_path=contract/protobuf/ {}
endif

swagger-gen:
	@which swag || (go install github.com/swaggo/swag/cmd/swag@latest)
	swag fmt
	swag init -g ../cmd/manager/main.go -d manager/ --parseDependency --output doc/swagger --instanceName manager
	swag init -g ../cmd/source/main.go -d source/ --parseDependency --output doc/swagger --instanceName source
	docker compose --env-file ./deployment/local/.env -f ./deployment/local/services/swagger.yml restart swagger

drop-manager-db:
	docker compose --env-file ./deployment/local/.env -f ./deployment/local/services/scylladb.yml exec scylladb cqlsh -e "DROP KEYSPACE manager;"