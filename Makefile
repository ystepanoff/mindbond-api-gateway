proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

api-gateway:
	go run cmd/main.go
