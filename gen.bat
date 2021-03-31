 protoc --proto_path=pbfiles --go_out=plugins=grpc:./../ pbfiles/*.proto

::gateway
 protoc --proto_path=pbfiles --grpc-gateway_out=logtostderr=true:./../ pbfiles/*.proto
