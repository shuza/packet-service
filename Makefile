include .env

protobuf:
	protoc -I . --go_out=plugins=micro:. ./proto/*.proto

docker_build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t shuzasa/packet-service:$(APP_VERSION) .

docker_run: docker_build
	docker run -p 8080:8080 -e BOX_SERVICE_ADDRESS=localhost:8081 \
	-e MONGO_HOST=192.168.99.100:30093 \
	-e USER_SERVICE_ADDRESS=$(USER_SERVICE_ADDRESS) \
	shuzasa/packet-service:$(APP_VERSION)

docker_push: docker_build
	docker push shuzasa/packet-service:$(APP_VERSION)

run:
	BOX_SERVICE_ADDRESS=$(BOX_SERVICE_ADDRESS) \
	MONGO_HOST=$(MONGO_HOST) \
	USER_SERVICE_ADDRESS=$(USER_SERVICE_ADDRESS) \
	go run main.go

deploy: docker_build
	kubectl apply -f ./deployment/packet-deployment.yaml

delete:
	kubectl delete svc packet-service
	kubectl delete deployment packet-deployment