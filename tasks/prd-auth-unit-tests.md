# PRD: Auth-Service Unit Tests for Error Mapping

## Introduction/Overview
Add unit tests across auth-service packages to validate error mapping, HTTP status/code consistency, and message safety. These tests will ensure all auth error paths return the correct `AppError` codes and messages without leaking internal details.

## Goals
- Verify error-code/status mappings across auth-service usecases and repositories.
- Ensure public messages are stable and safe for client responses.
- Catch regressions in error translation logic.
- Provide maintainable, table-driven tests for key auth flows.

## User Stories
- As a developer, I can refactor auth logic without breaking error contracts.
- As a maintainer, I can see exactly which error path maps to which HTTP status.
- As an API client, I consistently receive correct status codes and messages.

## Functional Requirements
1. Tests must cover error mapping in usecases, repositories, handlers, and response helpers where applicable.
2. Tests must assert both `AppError` codes and public messages.
3. Usecases must be tested with fakes/mocks to trigger each error branch.
4. Repository tests must validate translation of `gorm.ErrRecordNotFound` and other errors.
5. Test cases must be table-driven with clear names and expected outputs.
6. Table-driven tests must use `t.Run` with descriptive names for fast failure diagnosis.
7. Error messages in tests must include both expected and actual values for clarity.

## Non-Goals (Out of Scope)
- Full integration tests with real databases or external services.
- Performance or load testing.

## Design Considerations
- Prefer lightweight hand-written fakes for interfaces.
- Keep tests colocated with the code under test.
- Use helper functions to extract `AppError` for assertions.
- Use table-driven structures (slice or map) for test cases to reduce duplication.

## Technical Considerations
- Assert `GetErrCode()` and `PublicMessage()` for user-facing outputs.
- Use `errors.As` to validate `AppError` types.
- Avoid real JWT signing where not necessary by stubbing or using known tokens.
- Prefer `t.Run` per case; consider `t.Parallel()` if cases are independent.

## Success Metrics
- 100% of known error branches in auth-service are tested.
- Tests fail on incorrect error codes/messages.
- `go test ./auth-service/...` passes reliably.

## Open Questions
- Do we want to test handler-level response JSON, or just usecases/repositories?
- Should JWT parsing be unit-tested with real tokens or mocked?

## Plan
1. Inventory error branches in usecases, repositories, and handlers.
2. Add fakes for `AuthProfileRepo`, `RefreshTokenRepo`, and `UserGrpcClient`.
3. Write table-driven tests with `t.Run` names for `Register`, `Login`, `RefreshToken`, `ValidateToken`.
4. Use clear expected/actual error messages in assertions for fast diagnosis.
5. Add repository tests for `GetByEmail` and `GetByToken` error translation.
6. (Optional) Add handler/response tests for error-to-HTTP mapping.
