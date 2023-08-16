package utils

import "github.com/go-playground/validator/v10"

func ValidateStruct[T any](data T) error {
	validate := validator.New()
	err := validate.Struct(data)
	return err
}
