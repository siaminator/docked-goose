build:
	docker-compose run goose go build -o goose ./cmd

status:
	docker-compose up -d postgres
	docker-compose run goose ./goose status
