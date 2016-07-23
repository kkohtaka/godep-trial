# godep-trial

## How to build and run `server`

```bash
cd glide/server
glide install
go build
```

## How to update `.pb.go` files

```bash
cd glide/server
protoc -I ./pkg/hello/ ./pkg/hello/proto/*.proto --go_out=plugins=grpc:./pkg/hello
```
