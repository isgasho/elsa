@echo off

echo "start generate the grpc proto..."
 protoc -I ../pkg/proto   --go_out=plugins=grpc:../pkg/proto/pb  ../pkg/proto/*.proto
echo "generate the grpc code success..."