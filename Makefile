dependecies:
	docker-compose -f infra/compose.yaml up -d

build-app:
	go build cmd/main.go