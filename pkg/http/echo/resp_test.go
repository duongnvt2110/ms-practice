package echo

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	autherror "ms-practice/auth-service/pkg/utils/errors"
	apperror "ms-practice/pkg/errors"

	"github.com/labstack/echo/v4"
)

func TestResponseWithError_AppError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ResponseWithError(c, autherror.ErrInvalidToken)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resp APIResponse
	if decodeErr := json.Unmarshal(rec.Body.Bytes(), &resp); decodeErr != nil {
		t.Fatalf("failed to decode response: %v", decodeErr)
	}

	if rec.Code != autherror.ErrInvalidToken.GetHttpCode() {
		t.Fatalf("expected status %d, got %d", autherror.ErrInvalidToken.GetHttpCode(), rec.Code)
	}
	if resp.Message != autherror.ErrInvalidToken.PublicMessage() {
		t.Fatalf("expected message %q, got %q", autherror.ErrInvalidToken.PublicMessage(), resp.Message)
	}
	if resp.Code != autherror.ErrInvalidToken.GetErrCode() {
		t.Fatalf("expected code %s, got %s", autherror.ErrInvalidToken.GetErrCode(), resp.Code)
	}
}

func TestResponseWithError_NonAppError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := ResponseWithError(c, errors.New("boom"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resp APIResponse
	if decodeErr := json.Unmarshal(rec.Body.Bytes(), &resp); decodeErr != nil {
		t.Fatalf("failed to decode response: %v", decodeErr)
	}

	if rec.Code != apperror.ErrInternalServerError.GetHttpCode() {
		t.Fatalf("expected status %d, got %d", apperror.ErrInternalServerError.GetHttpCode(), rec.Code)
	}
	if resp.Message != apperror.ErrInternalServerError.PublicMessage() {
		t.Fatalf("expected message %q, got %q", apperror.ErrInternalServerError.PublicMessage(), resp.Message)
	}
	if resp.Code != apperror.ErrInternalServerError.GetErrCode() {
		t.Fatalf("expected code %s, got %s", apperror.ErrInternalServerError.GetErrCode(), resp.Code)
	}
}
