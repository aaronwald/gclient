```
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
go get -u google.golang.org/grpc
cd $GOHOME/src/
protoc --go_out=plugins=grpc:. -I /home/wald/dev/coypu/src/proto/ coincache.proto
```
