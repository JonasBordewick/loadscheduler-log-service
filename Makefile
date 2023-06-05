TAG ?= 0.0.1
DOCKER_USERNAME ?= jbordewick
DOCKER_REPOSITORY ?= loadscheduler-log-service

proto:
	protoc --go_out=. --go-grpc_out=. ./resources/*.proto

docker:
	docker build -t ${DOCKER_USERNAME}/${DOCKER_REPOSITORY}:${TAG} .
	docker push ${DOCKER_USERNAME}/${DOCKER_REPOSITORY}:${TAG}

release:
	docker pull ${DOCKER_USERNAME}/${DOCKER_REPOSITORY}:${TAG}
	docker tag ${DOCKER_USERNAME}/${DOCKER_REPOSITORY}:${TAG} ${DOCKER_USERNAME}/${DOCKER_REPOSITORY}:latest
	docker push ${DOCKER_USERNAME}/${DOCKER_REPOSITORY}:latest