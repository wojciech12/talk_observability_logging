SERVICE_NAME=order-manager
DOCKER_PROJECT_NAME=talk-observability-log

docker_build:
	docker build -t wojciech12/$(DOCKER_PROJECT_NAME) .

start:
	docker-compose -f docker-compose.yml \
		--project-name $(DOCKER_PROJECT_NAME) up -d

stop:
	docker-compose -f docker-compose.yml \
		--project-name $(DOCKER_PROJECT_NAME) stop

rebuild:
	docker-compose -f docker-compose.yml \
		--project-name $(DOCKER_PROJECT_NAME) build
