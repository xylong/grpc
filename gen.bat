 protoc --proto_path=protos --go_out=plugins=grpc:./../ protos/model/*.proto
 protoc --proto_path=protos --go_out=plugins=grpc:./../ protos/*.proto

::gateway
 protoc --proto_path=protos --grpc-gateway_out=logtostderr=true:./../ protos/model/*.proto
 protoc --proto_path=protos --grpc-gateway_out=logtostderr=true:./../ protos/*.proto
