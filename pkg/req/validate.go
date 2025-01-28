package req

import (
	"github.com/go-playground/validator/v10"
)

func Isvalid[T any](payload T) error {
	validate := validator.New()
	err := validate.Struct(payload)
	return err
}
