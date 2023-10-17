proto:
	git clone git@flotta-home:mindbond/proto.git
	protoc proto/auth.proto --go_out=plugins=grpc:./pkg/auth/
	protoc proto/chat.proto --go_out=plugins=grpc:./pkg/chat/
	rm -rf proto/

api-gateway:
	go run cmd/main.go
