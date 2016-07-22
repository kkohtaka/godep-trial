# godep-trial

## How to build and run `server`

```bash
cd server
protoc -I ./pkg/hello/ ./pkg/hello/proto/*.proto --go_out=plugins=grpc:./pkg/hello
go build
./server
```
