
build:
	go build -o proxy ./cmd


test:
	docker-compose -f ./deployments/docker-compose.yml up --build
	go test ./... -v
