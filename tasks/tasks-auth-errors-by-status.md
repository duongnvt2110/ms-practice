## Relevant Files

- `auth-service/pkg/usecases/auth_profile.go` - Auth usecase error mapping.
- `auth-service/pkg/handler/http/auth/token_auth.go` - Token-required error in logout handler.
- `auth-service/pkg/handler/http/http.go` - Auth middleware error mapping.
- `auth-service/pkg/repositories/auth_profile.go` - User lookup error mapping.
- `auth-service/pkg/repositories/refresh_token.go` - Refresh token lookup error mapping.
- `pkg/errors/errorlist.go` - Central error list and status codes.
- `pkg/http/echo/resp.go` - Response helper uses public messages.

### Notes

- Use `go test ./auth-service/...` to validate auth-service behavior after changes.

## Instructions for Completing Tasks

As you complete each task, check it off in this file by changing `- [ ]` to `- [x]`.

## Tasks

- [ ] 0.0 Create feature branch
  - [ ] 0.1 Create and checkout a new branch for this feature.
- [x] 1.0 Inventory auth-service error sites and map to HTTP statuses
  - [x] 1.1 Identify `errors.New(...)` in auth handlers, usecases, and repos.
  - [x] 1.2 Decide status mappings for each error message.
- [x] 2.0 Define/extend auth error list with missing status codes/messages
  - [x] 2.1 Add status-specific errors and codes to `pkg/errors/errorlist.go`.
- [x] 3.0 Replace plain errors with AppError mappings in auth flows
  - [x] 3.1 Replace usecase errors with `AppError` values or wraps.
  - [x] 3.2 Replace handler errors with `AppError` values.
  - [x] 3.3 Update repository error mapping for user lookup.
- [x] 4.0 Align response handling to use public messages only
  - [x] 4.1 Confirm response helper uses `PublicMessage()`.
- [ ] 5.0 Add tests for status mappings and key auth workflows
  - [ ] 5.1 Add unit tests for login/refresh/validate error status mapping.
  - [ ] 5.2 Add tests for user already exists and user not found flows.
