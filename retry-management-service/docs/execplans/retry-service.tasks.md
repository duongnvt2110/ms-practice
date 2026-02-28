# Retry Service ExecPlan Task Breakdown

This task breakdown derives from `docs/execplans/retry-service.md`. It is a planning artifact only and does not change code. Do not start implementation until the user explicitly approves.

## Phase 0: Prep

- [ ] Confirm required environment variables for retry-management-service (APP_HOST, APP_PORT, MYSQL_* vars, KAFKA_BROKERS, DLQ_TOPIC, DLQ_GROUP_ID).
- [ ] Verify MySQL database exists for the retry service and credentials work.
- [ ] Decide migration execution path (existing migration runner or manual execution).

## Phase 1: Configuration and Data Model

- [ ] Create retry-service config in `retry-management-service/pkg/config/config.go` using `godotenv` + `envconfig`.
- [ ] Add DLQ record model in `retry-management-service/pkg/model/dlq_record.go`.
- [ ] Add SQL migration in `retry-management-service/db/migrations/` to create `dlq_records` table.

## Phase 2: Persistence and Use Case

- [ ] Implement repository interface and GORM-backed repository in `retry-management-service/pkg/repository/dlq_repository.go`.
- [ ] Implement DLQ use case with JSON preview detection in `retry-management-service/pkg/usecase/dlq_usecase.go`.
- [ ] Update container wiring in `retry-management-service/pkg/container/container.go`.

## Phase 3: Kafka Consumer

- [ ] Add consumer runner in `retry-management-service/pkg/consumer/dlq_consumer.go`.
- [ ] Wire consumer startup in `retry-management-service/cmd/main.go` using `errgroup.WithContext`.
- [ ] Ensure offsets are committed after successful DB write.

## Phase 4: Admin UI (Server-Rendered)

- [ ] Implement HTTP server setup in `retry-management-service/pkg/handler/http/http.go`.
- [ ] Add routes in `retry-management-service/pkg/handler/http/routes.go`.
- [ ] Build admin handlers for list and detail in `retry-management-service/pkg/handler/http/admin`.
- [ ] Create templates under `retry-management-service/pkg/handler/http/admin/templates`.

## Phase 5: Validation

- [ ] Add a simple local Kafka produce helper (optional) or document how to produce a DLQ message.
- [ ] Run the service and verify ingestion and UI output.
- [ ] Capture expected logs and verification notes.
