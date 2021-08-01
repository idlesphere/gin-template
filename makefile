dotenv = "./config/.env"
compose_file = "./config/docker-compose.yml"

up:
	docker-compose --env-file $(dotenv) --file $(compose_file) up -d

down:
	docker-compose --env-file $(dotenv) --file $(compose_file) down