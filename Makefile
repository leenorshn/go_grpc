protos:
	protoc -I protos/ proto/currency.proto --go_out=plugins=grpc:proto/currency