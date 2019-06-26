include .env

protobuf:
	protoc -I . --go_out=plugins=grpc:. ./proto/*.proto

docker_build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shuzasa/packet-service:$(APP_VERSION) .

docker_run: docker_build
	docker run -p 8080:8080 -e BOX_SERVICE_ADDRESS=localhost:8081 \
	-e MONGO_HOST=192.168.99.100:30093 \
	shuzasa/packet-service:$(APP_VERSION)

docker_push: docker_build
	docker push shuzasa/packet-service:$(APP_VERSION)

run:
	PORT=$(PORT) \
	BOX_SERVICE_ADDRESS=$(BOX_SERVICE_ADDRESS) \
	MONGO_HOST=$(MONGO_HOST) \
	go run main.go