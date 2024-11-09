
build:
	go build -o go-balancer ./cmd


test:
	docker-compose -f ./deployments/docker-compose.yml up --build -d
	go test ./... -v
	docker-compose -f ./deployments/docker-compose.yml down
