# PRD: Auth-Service Error Grouping by HTTP Status

## Introduction/Overview
Standardize all `errors.New(...)` in auth-service by mapping them to explicit HTTP statuses using `AppError` so handlers return consistent status codes, public messages, and internal details for logs. This improves API consistency, observability, and maintainability.

## Goals
- Ensure all auth-service errors are mapped to clear HTTP statuses.
- Provide consistent public messages and stable error codes in responses.
- Keep internal details for logging and debugging (public vs internal).
- Reduce ad-hoc `errors.New(...)` usage and centralize error definitions.

## User Stories
- As an API client, I receive consistent status codes and error messages across auth endpoints.
- As a developer, I can reason about error handling without scanning every handler.
- As a maintainer, I can update error responses centrally.

## Functional Requirements
1. All `errors.New(...)` in auth-service must be replaced with `AppError` values (or wraps).
2. Each error must map to an explicit HTTP status (400/401/403/404/409/500 as needed).
3. Response helpers must send public-safe messages, not internal details.
4. Internal details must be preserved for logs (via wrapping).
5. Error codes must remain stable where possible, and new codes can be added when needed.

## Non-Goals (Out of Scope)
- Migrating other services to this error model.
- Introducing a global logging framework.
- Changing non-auth service response shapes.

## Design Considerations
- Centralize error definitions in `auth-service/pkg/utils/errors/errorlist.go`.
- Prefer `ErrBadRequest`, `ErrUnauthorized`, `ErrForbidden`, `ErrNotFound`, `ErrConflict`, `ErrInternalServer`.
- Provide a public message and optional internal detail (wrapped cause).

## Technical Considerations
- Use `errors.Is/As` for error detection in handlers and responses.
- Use `Wrap/Wrapf` when preserving internal error causes.
- Ensure error codes are unique and documented.

## Success Metrics
- 100% of auth-service error paths map to an `AppError`.
- No HTTP responses include internal error details.
- Unit tests verify status mapping for critical flows (login, refresh, validate token).

## Open Questions
- Should `403` vs `401` be used consistently for auth failures? (Proposed: `401` for missing/invalid token, `403` for valid token but insufficient permissions.)
- How should `409` be used for “user already exists” vs “bad request”?

## Plan (Auth Service Only)
1. Inventory all `errors.New(...)` in auth-service and map to HTTP statuses.
2. Extend `errorlist.go` with missing status-specific errors and codes.
3. Replace plain errors with `AppError` (or `Wrap/Wrapf`) in usecases/handlers.
4. Ensure response helpers use public messages only.
5. Add tests for status mapping and key endpoints.
