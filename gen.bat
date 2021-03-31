 protoc --proto_path=pbfiles --go_out=plugins=grpc:./../ prod.proto

::gateway
 protoc --proto_path=pbfiles --grpc-gateway_out=logtostderr=true:./../ prod.proto
