# connect-go-example

https://connect.build/docs/go/getting-started

```shell
mkdir connect-go-example
cd connect-go-example
go mod init example
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
```

define a service

```shell
mkdir -p greet/v1
touch greet/v1/greet.proto
```

generate code

```shell
buf mod init
```

create buf.gen.yaml

```shell
buf lint
buf generate
```