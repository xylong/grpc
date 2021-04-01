 protoc --proto_path=protos --go_out=plugins=grpc:./../ --validate_out=lang=go:./../ protos/model/*.proto
 protoc --proto_path=protos --go_out=plugins=grpc:./../ protos/*.proto

::gateway
 protoc --proto_path=protos --grpc-gateway_out=logtostderr=true:./../ protos/model/*.proto
 protoc --proto_path=protos --grpc-gateway_out=logtostderr=true:./../ protos/*.proto
