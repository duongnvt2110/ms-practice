package repositories

import (
	"context"
	"errors"
	"testing"
	"time"

	"ms-practice/auth-service/pkg/models"
	autherror "ms-practice/auth-service/pkg/utils/errors"
	apperror "ms-practice/pkg/errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newRefreshTokenTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	dsn := "file:refresh_token_test_" + time.Now().Format("20060102150405.000000000") + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}
	if err := db.AutoMigrate(&models.AuthProfile{}, &models.AuthRefreshToken{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func assertRepoAppError(t *testing.T, err error, expected apperror.AppError) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	var appErr apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T", err)
	}
	if appErr.GetErrCode() != expected.GetErrCode() {
		t.Fatalf("expected err code %s, got %s", expected.GetErrCode(), appErr.GetErrCode())
	}
	if appErr.PublicMessage() != expected.PublicMessage() {
		t.Fatalf("expected message %q, got %q", expected.PublicMessage(), appErr.PublicMessage())
	}
}

func TestRefreshTokenRepo_GetByToken_NotFound(t *testing.T) {
	db := newRefreshTokenTestDB(t)
	repo := NewRefreshTokenRepo(db)

	_, err := repo.GetByToken(context.Background(), "missing")
	assertRepoAppError(t, err, autherror.ErrInvalidRefreshToken)
}

func TestRefreshTokenRepo_GetByToken_InternalError(t *testing.T) {
	db := newRefreshTokenTestDB(t)
	if err := db.Migrator().DropTable(&models.AuthRefreshToken{}); err != nil {
		t.Fatalf("failed to drop table: %v", err)
	}
	repo := NewRefreshTokenRepo(db)

	_, err := repo.GetByToken(context.Background(), "missing")
	assertRepoAppError(t, err, apperror.ErrInternalServerError)
}
