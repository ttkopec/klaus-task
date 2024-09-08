proto:
	protoc --proto_path=./pkg/protos --go_out=./pkg/protos --go-grpc_out=./pkg/protos ./pkg/protos/score_service.proto

run:
	DEBUG=1 go run cmd/app.go

.PHONY: mocks
mocks: 
	go install github.com/vektra/mockery/v2@v2.45.0
	mockery

test:
	go test ./... -v

docker_build:
	docker buildx build --platform linux/arm64 -t klaus-task .

docker_run:
	docker run --rm --publish 5001:5001 klaus-task