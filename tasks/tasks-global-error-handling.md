## Relevant Files

- `auth-service/pkg/utils/errors/errors.go` - Core auth-service error type and wrapping behavior.
- `auth-service/pkg/utils/errors/stack.go` - Stack trace capture/formatting helpers.
- `auth-service/pkg/utils/errors/errors_test.go` - Unit tests for wrapping, stack formatting, and public messages.
- `auth-service/pkg/utils/errors/errorlist.go` - Auth-service error definitions (codes/messages).
- `pkg/http/echo/resp.go` - Shared response helper used by auth-service.

### Notes

- Unit tests should live next to the code under `auth-service/pkg/utils/errors/`.
- Use `go test ./auth-service/pkg/utils/errors` to run the error package tests.

## Instructions for Completing Tasks

As you complete each task, check it off in this file by changing `- [ ]` to `- [x]`.

## Tasks

- [ ] 0.0 Create feature branch
  - [ ] 0.1 Create and checkout a new branch for this feature.
- [x] 1.0 Review current auth-service error package usage and touchpoints
  - [x] 1.1 Identify auth-service handlers and response helpers using AppError.
  - [x] 1.2 Identify any Wrap/Wrapf usage within auth-service.
- [x] 2.0 Define the new error API and stack-trace behavior (auth-service only)
  - [x] 2.1 Confirm public message vs internal detail behavior.
  - [x] 2.2 Decide on error chain behavior (`Unwrap`, `errors.Is/As`).
  - [x] 2.3 Specify `%+v` output expectations (message + stack + cause).
- [x] 3.0 Implement the error package changes in auth-service
  - [x] 3.1 Update error type fields and interface methods.
  - [x] 3.2 Implement `Unwrap` and `Format` for stack traces.
  - [x] 3.3 Update `Wrap`/`Wrapf` to preserve cause chaining.
  - [x] 3.4 Update `Catch` to use `errors.As`.
- [x] 4.0 Update auth-service response handling to use public messages
  - [x] 4.1 Update `pkg/http/echo/resp.go` to import the correct package.
  - [x] 4.2 Use public message when building HTTP error responses.
  - [x] 4.3 Use `errors.As` for error detection in responses.
- [x] 5.0 Add tests and minimal docs for wrapping/stack behavior
  - [x] 5.1 Add unit tests for `Wrap` and `errors.Is/As` behavior.
  - [x] 5.2 Add unit test for `%+v` to include stack/cause.
  - [x] 5.3 Add unit test for public message vs detail behavior.
- [x] 6.0 Optimize stack formatting and helper utilities
  - [x] 6.1 Adjust `Stack.Format` to emit full frames only for `%+v`.
  - [x] 6.2 Make `Stack.String()` return a compact summary.
  - [x] 6.3 Add a `WithStack` helper for external errors.
  - [x] 6.4 Add tests for compact `%v/%s` output.
