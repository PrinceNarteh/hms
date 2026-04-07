.PHONY: start-dev
start-dev:
	@docker-compose -f docker-compose.dev.yml up --build

.PHONY: stop-dev
stop-dev:
	@docker-compose -f docker-compose.dev.yml down
