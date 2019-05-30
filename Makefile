protobuf:
	protoc -I . --go_out=plugins=grpc:. ./proto/*.proto

build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shuzasa/packet-service .

run:
	docker run -p 8000:8000 shuzasa/packet-service