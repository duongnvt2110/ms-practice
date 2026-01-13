package errors

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestWrapPreservesCauseChain(t *testing.T) {
	root := errors.New("root cause")
	err := ErrInternalServerError.Wrap(root)

	if !errors.Is(err, root) {
		t.Fatalf("expected errors.Is to match root cause")
	}
	if err.PublicMessage() != "internal server error" {
		t.Fatalf("expected public message to remain unchanged")
	}
}

func TestWrapfAddsDetailAndFormatsStack(t *testing.T) {
	root := errors.New("db down")
	err := ErrInternalServerError.Wrapf(root, "while saving user %d", 42)

	out := fmt.Sprintf("%+v", err)
	if !strings.Contains(out, "internal server error") {
		t.Fatalf("expected output to include public message")
	}
	if !strings.Contains(out, "while saving user 42") {
		t.Fatalf("expected output to include detail message")
	}
	if !strings.Contains(out, "Caused by:") {
		t.Fatalf("expected output to include cause marker")
	}
	if !strings.Contains(out, "errors_test.go") {
		t.Fatalf("expected output to include stack trace")
	}
}

func TestCatchUsesErrorsAs(t *testing.T) {
	base := ErrUnauthorized.Wrap(errors.New("bad token"))
	wrapped := fmt.Errorf("outer: %w", base)

	caught := ErrInternalServerError.Catch(wrapped)
	if caught != base {
		t.Fatalf("expected Catch to return the wrapped AppError")
	}
}

func TestStackFormattingIsCompactByDefault(t *testing.T) {
	stack := StackTrace(0)

	compact := fmt.Sprintf("%v", stack)
	if strings.Contains(compact, "\n") {
		t.Fatalf("expected compact stack output without newlines")
	}

	verbose := fmt.Sprintf("%+v", stack)
	if !strings.Contains(verbose, "errors_test.go") {
		t.Fatalf("expected verbose stack output to include file name")
	}
	if !strings.Contains(verbose, "\n") {
		t.Fatalf("expected verbose stack output to include newlines")
	}
}

func TestWithStackAddsStackTrace(t *testing.T) {
	err := ErrInternalServerError.Wrap(errors.New("boom"))
	out := fmt.Sprintf("%+v", err)
	if !strings.Contains(out, "boom") {
		t.Fatalf("expected output to include error message")
	}
	if !strings.Contains(out, "errors_test.go") {
		t.Fatalf("expected output to include stack trace")
	}
}
