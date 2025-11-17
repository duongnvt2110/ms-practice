DOCKER_COMPOSE ?= docker compose
SERVICES ?= api-gateway-service auth-service user-service booking-service payment-service catalog-service ticket-service
COMPOSE_FILES := $(addsuffix /build/docker-compose.yaml,$(SERVICES))

.PHONY: up down restart up-% down-% logs-% status

up:
	@$(foreach file,$(COMPOSE_FILES), \
		echo "Starting services defined in $(file)"; \
		$(DOCKER_COMPOSE) -f $(file) up -d --build || exit $$?; )

down:
	@$(foreach file,$(COMPOSE_FILES), \
		if [ -f $(file) ]; then \
			echo "Stopping services defined in $(file)"; \
			$(DOCKER_COMPOSE) -f $(file) down --remove-orphans; \
		fi; )

restart: down up

up-%:
	$(DOCKER_COMPOSE) -f $*/build/docker-compose.yaml up -d --build

down-%:
	$(DOCKER_COMPOSE) -f $*/build/docker-compose.yaml down --remove-orphans

logs-%:
	$(DOCKER_COMPOSE) -f $*/build/docker-compose.yaml logs -f

status:
	@$(foreach file,$(COMPOSE_FILES), \
		if [ -f $(file) ]; then \
			echo "Status for $(file):"; \
			$(DOCKER_COMPOSE) -f $(file) ps; \
			echo ""; \
		fi; )
