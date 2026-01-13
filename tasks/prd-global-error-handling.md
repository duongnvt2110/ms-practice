# PRD: Global Error Handling (Auth Service First)

## Introduction/Overview
Standardize and improve the `auth-service` error library to support richer diagnostics (stack traces and causes) while defining a reusable pattern that can later become the global solution across services. The immediate scope is auth-service only, but the design must be portable.

## Goals
- Provide consistent error wrapping with stack traces and root cause chaining.
- Preserve a clear public message for HTTP responses while retaining internal details.
- Define a portable API that can be adopted by other services later.
- Reduce ambiguity in error handling by aligning with Go `errors.Is/As` behavior.
- Keep stack trace output compact by default and verbose only when explicitly requested (`%+v`).

## User Stories
- As a backend developer, I can wrap errors and retain original causes for debugging.
- As an API client, I receive a stable public error message and code.
- As a maintainer, I can apply the same error package to other services with minimal changes.

## Functional Requirements
1. The error type must support wrapping a cause such that `errors.Is/As` work as expected.
2. The error type must capture a stack trace at wrap time, and print it only on `%+v` formatting.
3. The error type must expose a public message and a code for HTTP responses.
4. The error type must retain an internal error/cause for logging and debugging.
5. The error package must be self-contained in `auth-service` and not require other services to build.
6. The response layer must use the public message and code, not the internal cause.
7. The stack trace string representation should be compact by default (no full trace on `%s`/`%v`).
8. Provide a helper (e.g., `WithStack`) to attach a stack trace to external errors when needed.

## Non-Goals (Out of Scope)
- Migrating other services to the new package in this phase.
- Changing external HTTP contracts for auth-service beyond message/code semantics.
- Introducing a centralized logging framework.

## Design Considerations
- Keep the API aligned with existing usage (`Wrap`, `Wrapf`, `GetHttpCode`, `GetErrCode`) to minimize refactors in auth-service.
- The public message must be stable and safe for client exposure.

## Technical Considerations
- Use `errors.Is/As` and `Unwrap()` on the new error type.
- Implement `Format(fmt.State, rune)` to display the stack trace only when using `%+v`.
- Keep `%s`/`%v` output compact; `String()` should not emit full stack traces.
- Capture stacks using `runtime.Callers` with a consistent skip depth.
- Provide a `WithStack` helper to attach stack traces to external errors.
- Follow the DoltHub pattern for stack traces and error formatting.

## Success Metrics
- 100% of new auth-service error wraps preserve cause chain (`errors.Is/As` pass in tests).
- Stack traces appear only for `%+v` in logs (verified by unit test output).
- `%s`/`%v` output remains compact (verified by unit test output).
- No client-facing errors leak internal error details.

## Open Questions
- Do we want to add a structured field (e.g., `details`) to the JSON error response later?
- Should error codes be globally unique across services or only within a service?

## Example API Sketch (Auth Service Only)
```go
// Public interface remains similar.
type AppError interface {
    error
    Unwrap() error
    WithStack(err error) AppError
    Wrap(cause error) AppError
    Wrapf(cause error, format string, a ...interface{}) AppError
    GetHttpCode() int
    GetErrCode() string
    PublicMessage() string
}
```

```go
// Example usage in a handler.
if err := svc.Do(ctx); err != nil {
    appErr := errors.ErrInternalServer.Wrap(err)
    response.Error(c, appErr) // sends public message + code
    log.Printf("%+v", appErr) // includes stack trace + cause
}
```

## Plan (Auth Service First)
1. Define the new error type in `auth-service/pkg/utils/errors` with `Unwrap` and `Format`.
2. Update `Wrap/Wrapf` to wrap the cause in the error chain and record stack.
3. Add `PublicMessage()` or equivalent to separate public and internal messages.
4. Update auth-service response helpers to use the public message.
5. Add unit tests for wrapping, `errors.Is/As`, and `%+v` stack trace output.
6. Document the intended path for adoption by other services.
