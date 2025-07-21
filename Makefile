run:
	docker-compose up --build

down:
	docker-compose down

logs:
	docker-compose logs -f app

ps:
	docker-compose ps

migrate:
	docker exec -it fashiondemon_app ./server migrate

fmt:
	go fmt ./...

test:
	go test ./...
