# Retry DLQ Ingestion and Admin Review UI

This ExecPlan is a living document. The sections `Progress`, `Surprises & Discoveries`, `Decision Log`, and `Outcomes & Retrospective` must be kept up to date as work proceeds. This plan follows the ExecPlan requirements in `/Users/exe-macbook/duong/dotfiles/dotfiles/.codex/.agent/PLANS.md`.

## Purpose / Big Picture

After this work, the retry-management-service will continuously consume failed messages from the Kafka dead-letter topic `dlq.topics`, store each message in MySQL, and expose a server-rendered admin web UI to review the stored records. A developer or admin will be able to run the service, open a browser to a local admin page, and see a list of DLQ records and a detail view that includes message metadata and payload.

## Progress

- [x] (2026-02-10 03:49Z) Capture requirements and create the initial ExecPlan.
- [x] (2026-02-10 04:15Z) Implement configuration, model, migration, repository, use case, HTTP handlers, templates, and consumer wiring.
- [ ] Validate end-to-end ingestion and UI review locally.

## Surprises & Discoveries

- None so far.

## Decision Log

- Decision: Treat `dlq.topics` as a single Kafka topic name, not a prefix or list.
  Rationale: The user provided a single topic identifier without delimiters or prefix semantics.
  Date/Author: 2026-02-10 / Codex
- Decision: Store raw message payload and headers in the database and render a best-effort JSON preview in the UI when the payload is valid JSON.
  Rationale: No message schema was provided, so raw storage preserves all data while still enabling basic human review.
  Date/Author: 2026-02-10 / Codex
- Decision: Use a dedicated Kafka reader in the retry service to commit offsets only after a successful database write.
  Rationale: The shared Kafka helper commits on read, which does not satisfy the post-write commit requirement.
  Date/Author: 2026-02-10 / Codex

## Outcomes & Retrospective

- Not started. This section will be updated after implementation milestones complete.

## Context and Orientation

This repository is a Go 1.23.6 monorepo with multiple microservices. The retry-management-service currently contains only a basic folder skeleton created to mirror `user-service`. The service will be extended to run an HTTP server using `github.com/gorilla/mux` (already a dependency in the repo) and a Kafka consumer using `github.com/segmentio/kafka-go` (already a dependency). MySQL access is provided by the shared GORM client at `pkg/db/gorm_client` in the repository root, which accepts a `pkg/config.Mysql` configuration. A “dead-letter queue” or “DLQ” here means a Kafka topic that receives messages that could not be processed successfully by upstream services. The plan will add a database table to store each DLQ message with metadata (topic, partition, offset, key, headers, timestamps) and raw payload bytes.

Key existing files and folders to orient a new reader:

In `user-service`, `user-service/pkg/handler/http/http.go` shows the pattern for a gorilla/mux HTTP server with graceful shutdown; the retry service should follow that structure. The shared MySQL configuration struct is in `pkg/config/config.go` at repo root. The Kafka helper is in `pkg/kafka/kafka.go` at repo root; it can be used for consuming, or we can instantiate a `kafka.Reader` directly if needed, but the plan below will reuse the helper for consistency.

## Plan of Work

Milestone 1 focuses on service configuration, data model, and database storage. We will add a retry-service config in `retry-management-service/pkg/config` that wraps the shared `pkg/config.App`, `pkg/config.Mysql`, and `pkg/config.Kafka` configuration, plus service-specific fields such as the DLQ topic name and consumer group ID. We will add a GORM model for a DLQ record, a repository that can insert records and query them by pagination and by ID, and a migration SQL file in `retry-management-service/db/migrations` to create the table. This milestone ends with a repository method that can write and read records in isolation.

Milestone 2 adds the Kafka consumer and use case logic. We will create a use case that accepts a `kafka.Message`, normalizes headers, and persists the record using the repository. The consumer will run in a background goroutine when the service starts, use a dedicated consumer group ID, and only commit offsets after a successful database write. For errors, we will log and continue; the goal is to keep the consumer alive and provide visibility in logs. This milestone ends with a running process that consumes from `dlq.topics` and stores records.

Milestone 3 delivers the server-rendered admin UI. We will create HTTP handlers that render an HTML list page (table of DLQ records with pagination) and a detail page. The UI will be rendered via Go `html/template` templates stored under `retry-management-service/pkg/handler/http/admin/templates`. The list page will show topic, partition, offset, key (truncated), created time, and a link to details. The detail page will show full metadata, headers, and a payload preview. This milestone ends with an HTTP server serving `/admin/dlq` and `/admin/dlq/{id}` with a basic, readable layout and no client-side framework.

## Concrete Steps

1. Add configuration for the retry service in `retry-management-service/pkg/config/config.go`. It should embed or include `pkg/config.App`, `pkg/config.Mysql`, and `pkg/config.Kafka`, and add `DlqTopic` and `DlqGroupID` fields with environment variables `DLQ_TOPIC` (default `dlq.topics`) and `DLQ_GROUP_ID` (default `retry-management-service`). Reuse the `godotenv` and `envconfig` pattern used in `user-service/pkg/config/config.go`.

2. Define the GORM model at `retry-management-service/pkg/model/dlq_record.go`. The model should include fields: `ID` (auto-increment), `Topic`, `Partition`, `Offset`, `Key`, `Headers` (JSON string), `Payload` (raw bytes), `PayloadJSON` (nullable text containing a best-effort pretty JSON string), `CreatedAt`. The model should be explicit about field types to produce a stable schema.

3. Create a migration SQL file in `retry-management-service/db/migrations/` that creates a table (for example, `dlq_records`) with the above fields and suitable types (`BIGINT` for offset, `LONGBLOB` for payload, `TEXT` for headers). Include `-- migrate:up` and `-- migrate:down` blocks to mirror other services.

4. Implement a repository in `retry-management-service/pkg/repository/dlq_repository.go` using GORM. It should expose `Create(ctx, record)` and `List(ctx, page, pageSize)` and `GetByID(ctx, id)` methods. Pagination should be deterministic: order by `id` descending.

5. Implement a use case in `retry-management-service/pkg/usecase/dlq_usecase.go` that accepts a `kafka.Message`, maps fields into a model, performs JSON detection (try to `json.Unmarshal` into `interface{}` and re-marshal with indentation for `PayloadJSON`), and calls the repository `Create`.

6. Wire dependencies in `retry-management-service/pkg/container/container.go`. The container should initialize config, the GORM client via `pkg/db/gorm_client`, the repository, and the use case. It should also build a Kafka client using `pkg/kafka.NewKafkaClient(cfg.Kafka)` and set the reader topic to `DlqTopic` and group ID `DlqGroupID`.

7. Implement the Kafka consumer runner in `retry-management-service/pkg/consumer/dlq_consumer.go` (new package). It should accept context, Kafka client, and use case, call `Consume` with a handler that writes to the DB. It should log errors and continue.

8. Implement the HTTP server in `retry-management-service/pkg/handler/http/http.go` and routes in `retry-management-service/pkg/handler/http/routes.go` following `user-service` patterns. Add handlers under `retry-management-service/pkg/handler/http/admin` for list and detail pages. Use `html/template` to render the templates.

9. Update `retry-management-service/cmd/main.go` to start both the HTTP server and the consumer, using `errgroup.WithContext` similar to `user-service/cmd/main.go`. The process should exit gracefully on interrupt.

## Validation and Acceptance

Run the service from the repository root with:

    go run ./retry-management-service/cmd

Then produce a test DLQ message using a local Kafka producer or `kafka-go` snippet (documented in the artifacts section). After producing at least one message to `dlq.topics`, verify:

1. The service logs show it started the HTTP server and the Kafka consumer without errors.
2. The MySQL table `dlq_records` contains a row with the correct topic, partition, offset, and payload.
3. Visiting `http://localhost:<APP_PORT>/admin/dlq` shows the record in the list.
4. Clicking the record opens `http://localhost:<APP_PORT>/admin/dlq/{id}` and displays metadata and payload preview.

Acceptance is met when a newly produced DLQ message appears in the admin UI within a few seconds and is persisted in the database.

## Idempotence and Recovery

All steps are additive and can be re-run safely. If the migration already exists, create a new timestamped migration file rather than editing existing ones. If the Kafka consumer fails to parse a payload as JSON, it should still store the raw payload and leave `PayloadJSON` empty, allowing ingestion to proceed. If the HTTP server fails to bind to the port, change `APP_PORT` in the environment and retry.

## Artifacts and Notes

Expected log line when the service starts:

    Server is running on http://<APP_HOST>:<APP_PORT>

Example expected SQL row after ingestion:

    id=1 topic=dlq.topics partition=0 offset=42 key=<...> created_at=<timestamp>

Example local test producer snippet (to be documented in the repo if needed):

    go run ./scripts/kafka-produce-dlq.go

## Interfaces and Dependencies

The retry service will depend on the following internal modules:

- `ms-practice/pkg/config`: Provides shared `App`, `Mysql`, and `Kafka` configuration structs and environment variable naming.
- `ms-practice/pkg/db/gorm_client`: Provides `NewGormClient(mysqlCfg config.Mysql) (*gorm.DB, error)` for MySQL connections.
- `ms-practice/pkg/kafka`: Provides `KafkaClient` with `SetReaderTopic(topic, groupId)` and `Consume(ctx, handler)` for Kafka consumption.

New or updated interfaces to define:

In `retry-management-service/pkg/repository/dlq_repository.go`, define:

    type DLQRepository interface {
        Create(ctx context.Context, record *model.DLQRecord) error
        List(ctx context.Context, page, pageSize int) ([]model.DLQRecord, int64, error)
        GetByID(ctx context.Context, id int64) (*model.DLQRecord, error)
    }

In `retry-management-service/pkg/usecase/dlq_usecase.go`, define:

    type DLQUsecase interface {
        Ingest(ctx context.Context, msg kafka.Message) error
        List(ctx context.Context, page, pageSize int) ([]model.DLQRecord, int64, error)
        GetByID(ctx context.Context, id int64) (*model.DLQRecord, error)
    }

The HTTP handlers will call the use case methods for list and detail. The Kafka consumer will call `Ingest`.

Change log:

This ExecPlan was created to meet the user request for a retry-management-service that consumes Kafka DLQ messages, stores them in MySQL, and exposes a server-rendered admin review UI.
