proto:
	protoc --proto_path=pkg/ --go_out=./ --go-grpc_out=./ pkg/pb/auth.proto

server:
	go run cmd/main.go
cover:
	go test -short -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out