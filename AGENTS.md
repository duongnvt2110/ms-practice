# Repository Guidelines

## Project Structure & Module Organization
Each microservice lives in its own directory (`auth-service`, `booking-service`, `payment-service`, etc.) and follows a consistent layout: `cmd/main.go` contains the HTTP or gRPC entrypoint, `pkg` holds service-specific business logic, and `build/docker-compose.yaml` describes how to run that service with its dependencies. Shared code such as Kafka clients, DTOs, and middleware belongs under `pkg` and `utils-service`, while `proto/` stores canonical protobuf definitions used across services. Documentation and architectural references live in `docs/` and the root diagrams (`architecture.excalidraw`, `archtecture.png`).

## Build, Test, and Development Commands
Use the root `Makefile` to coordinate environments: `make up` builds and starts every compose file under each service’s `build/` folder, and `make up-auth-service` (or any service name) runs just that stack for faster iteration. Tear everything down with `make down` or `make down-user-service`. When iterating on a single Go binary, run it directly (e.g., `go run ./auth-service/cmd`). gRPC stubs stay current by running `protoc --go_out --go-grpc_out ./proto/user.proto` (or `payment.proto`) exactly as shown in the README snippet.

## Coding Style & Naming Conventions
Target Go 1.23.6 and keep modules tidy with `go mod tidy` when dependencies change. Run `gofmt -w` (or `go fmt ./...`) before committing; prefer snake_case for environment variables, UpperCamelCase for exported structs/interfaces, and lowerCamelCase for private members. HTTP handler files typically mirror the resource (`user_handler.go`, `booking_routes.go`) inside the relevant service’s `pkg` folder. Keep Kafka topics and event structs singular and suffixed with `Event` to match the saga documentation.

## Testing Guidelines
Unit tests should mirror the package path (e.g., `auth-service/pkg/token/token_test.go`) and follow the `TestXxx` naming that Go’s `testing` package requires. Run `go test ./...` at the repo root to verify all services, and use `go test ./booking-service/... -run TestBookingSaga` to target a specific workflow. Aim for meaningful coverage on domain packages (models, repositories, handlers) and mock Kafka or gRPC clients using `go test` table-driven cases.

## Commit & Pull Request Guidelines
Existing commits use short, imperative subjects (“update payment service”, “fix repo and usecase”); keep new messages under ~72 characters and describe what the change does, not how. Every pull request should link back to the relevant requirement or issue, summarize affected services, and include the commands used for validation (`go test ./...`, `make up-user-service`). Add screenshots or curl examples when touching HTTP contracts to ease verification.

## Proto & Messaging Notes
Whenever protobufs change, regenerate code before opening a PR and mention the affected services. Keep Kafka topic definitions synchronized with those in `README.md`, and document any new event type in both the README saga section and the service README under `docs/`.


# Codex Guildlines 
- Do not implement any things before I approve