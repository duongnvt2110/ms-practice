package validator

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validateOnce sync.Once
	validate     *validator.Validate
)

func NewValidate() *validator.Validate {
	validateOnce.Do(func() {
		validate = validator.New()
	})
	return validate
}
