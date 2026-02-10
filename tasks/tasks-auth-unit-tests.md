## Relevant Files

- `auth-service/pkg/usecases/auth_profile_test.go` - Usecase unit tests for error mapping.
- `auth-service/pkg/usecases/auth_profile.go` - Usecase logic under test.
- `auth-service/pkg/utils/errors/errorlist.go` - Auth-specific error definitions used in tests.
- `pkg/errors/errorlist.go` - Common error definitions for fallback cases.
- `auth-service/pkg/repositories/auth_profile_test.go` - Repository tests for user lookup error mapping.
- `auth-service/pkg/repositories/refresh_token_test.go` - Repository tests for refresh token error mapping.
- `pkg/http/echo/resp_test.go` - Handler response mapping tests.

### Notes

- Use `go test ./auth-service/...` to run auth-service tests.

## Instructions for Completing Tasks

As you complete each task, check it off in this file by changing `- [ ]` to `- [x]`.

## Tasks

- [ ] 0.0 Create feature branch
  - [ ] 0.1 Create and checkout a new branch for this feature.
- [x] 1.0 Inventory auth-service error branches for testing
  - [x] 1.1 List error branches in `Register`, `Login`, `RefreshToken`, `ValidateToken`.
- [x] 2.0 Build fakes/mocks for auth-service dependencies
  - [x] 2.1 Add fakes for `AuthProfileRepo`, `RefreshTokenRepo`, and `UserGrpcClient`.
- [x] 3.0 Write usecase unit tests for error mappings
  - [x] 3.1 Add tests for Register, Login, RefreshToken, ValidateToken error cases.
  - [x] 3.2 Refactor tests to table-driven form with `t.Run` names.
  - [x] 3.3 Improve assertions to include expected vs actual values.
- [ ] 4.0 Write repository unit tests for error translation
  - [x] 4.1 Add tests for `GetByEmail` and `GetByToken` error mapping.
- [ ] 5.0 (Optional) Add handler/response tests for error-to-HTTP mapping
  - [x] 5.1 Add tests for response payloads and status codes.
- [ ] 6.0 Run test suite and document coverage gaps
  - [ ] 6.1 Run `go test ./auth-service/...` and note any failures.
