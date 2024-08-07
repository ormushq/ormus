package ormus

//go:generate sh -c "find . -regex '.*/contract/protobuf/.*/[^/]*.proto' -print0 | xargs -n 1 protoc -I contract/protobuf --go_out=./contract/go --go_opt=paths=source_relative --go-grpc_out=./contract/go --go-grpc_opt=paths=source_relative"
