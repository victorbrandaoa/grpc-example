# grpc-example

Dentro de server rodar:

```sh
protoc -I=.. --go_out=.. --go-grpc_out=.. server/proto/user/user.proto

go run main.go
```

Dentro de client rodar:

```sh
protoc -I=.. --go_out=.. --go-grpc_out=.. client/proto/user/user.proto

go run main.go
```
