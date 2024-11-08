
build:
	go build -o proxy ./cmd

docker_build:
	docker build -f ./deployments/Dockerfile .