DOCKER_COMPOSE ?= docker compose
DOCKER_SYSTEM_PRUNE ?= docker system prune

SERVICES := \
	api-gateway-service \
	auth-service \
	user-service \
	booking-service \
	payment-service \
	event-service \
	ticket-service \
	utils-service

# COMPOSE_api-gateway-service := api-gateway-service/build/docker-compose.yaml
COMPOSE_auth-service := auth-service/build/docker-compose.yaml
COMPOSE_user-service := user-service/build/docker-compose.yaml
COMPOSE_booking-service := booking-service/build/docker-compose.yaml
COMPOSE_payment-service := payment-service/build/docker-compose.yaml
COMPOSE_event-service := event-service/build/docker-compose.yaml
COMPOSE_ticket-service := ticket-service/build/docker-compose.yaml
COMPOSE_retry-management-service := retry-management-service/build/docker-compose.yaml
COMPOSE_utils-service := utils-service/docker-compose.yaml

COMPOSE_VARS := $(filter COMPOSE_%,$(.VARIABLES))

SERVICES := $(patsubst COMPOSE_%,%,$(COMPOSE_VARS))

COMPOSE_FILES := $(foreach service,$(SERVICES),$(COMPOSE_$(service)))

.PHONY: up down restart up-% down-% logs-% status

up:
	@$(foreach file,$(COMPOSE_FILES), \
		echo "Starting services defined in $(file)"; \
		$(DOCKER_COMPOSE) -f $(file) up -d --build || exit $$?; ) \
		$(DOCKER_SYSTEM_PRUNE)

down:
	@$(foreach file,$(COMPOSE_FILES), \
		if [ -f $(file) ]; then \
			echo "Stopping services defined in $(file)"; \
			$(DOCKER_COMPOSE) -f $(file) down --remove-orphans; \
		fi; )

restart: down up

up-%:
	$(eval COMPOSE_FILE := $(COMPOSE_$*))
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) up -d --build
	$(DOCKER_SYSTEM_PRUNE)

down-%:
	$(eval COMPOSE_FILE := $(COMPOSE_$*))
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) down --remove-orphans
	$(DOCKER_SYSTEM_PRUNE)

logs-%:
	$(eval COMPOSE_FILE := $(COMPOSE_$*))
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) logs -f

status:
	@$(foreach file,$(COMPOSE_FILES), \
		if [ -f $(file) ]; then \
			echo "Status for $(file):"; \
			$(DOCKER_COMPOSE) -f $(file) ps; \
			echo ""; \
		fi; )

exec-%:
	$(eval COMPOSE_FILE := $(COMPOSE_$*))
	$(eval SERVICE := $*)
	$(DOCKER_COMPOSE) -f $(COMPOSE_FILE) exec $(SERVICE) sh
