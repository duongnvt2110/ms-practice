# Notification Service

Provides email, push, and in-app notifications triggered from Kafka events such as `PaymentSucceeded` and `TicketIssued`.

## Responsibilities
- Consume `payments.events` + `tickets.events`
- Persist per-channel notification jobs in MySQL (`notification_jobs`)
- Run workers per channel (email/push/in-app) with retry + backoff
- Expose `/health` endpoint for readiness checks

## Database Tables
### notification_jobs
| column | type | note |
| --- | --- | --- |
| id | bigint | primary key |
| idempotency_key | varchar | unique (`event_type + event_id + channel`) |
| event_type | varchar |
| event_id | int |
| user_id | int |
| channel | enum(email,push,in_app) |
| template | varchar |
| payload | json |
| status | enum(pending,sending,sent,failed) |
| attempts | int |
| max_attempts | int |
| last_error | text |
| next_attempt_at | datetime |
| sent_at | datetime nullable |
| created_at/updated_at | datetime |

### notification_templates *(optional future enhancement)*
Stores template metadata to render notifications. Current implementation maps template codes in code but table is reserved for future UI-driven management.

## Environment Variables
| env | default | description |
| --- | --- | --- |
| APP_HOST | localhost | HTTP bind host |
| APP_PORT | 8006 | HTTP bind port |
| MYSQL_* | | DB connection (see shared config) |
| KAFKA_BROKERS | host.docker.internal:29092 | Kafka cluster |
| NOTI_MAX_ATTEMPTS | 3 | max retries per job |
| NOTI_RETRY_INTERVAL_SECONDS | 60 | backoff seconds |
| NOTI_WORKER_INTERVAL_MS | 500 | worker polling interval |
| SMTP_HOST/PORT/USERNAME/PASSWORD/FROM | | email provider settings |
| FIREBASE_* | | push provider credentials (stub) |

## Sequence
1. Payment service emits `PaymentSucceeded` payload.
2. Notification consumer persists 3 jobs (email/push/in-app) with shared payload.
3. Channel workers poll pending jobs, invoke provider, update status.
4. Failed jobs retry until `max_attempts` observed.

## Future Work
- Integrate real SMTP + Firebase clients.
- Fetch user contact preferences from User Service before enqueuing jobs.
- Implement admin APIs to list/retry jobs and manage templates.
