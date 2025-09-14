dependecies:
	docker-compose -f infra/compose.yaml up -d

build-app: dependecies
	go build cmd/main.go