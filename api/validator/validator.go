package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Client = *newValidatorClient()

var (
	ErrInvalidType  = errors.New("Error: you input invalid type, this type can't not be support ex) nil")
	ErrInvalidation = errors.New("Error: you're struct type not valid, please check strcut field")
)

type validatorClient struct {
	validator *validator.Validate
}

func newValidatorClient() *validatorClient {
	v := &validatorClient{
		validator: validator.New(),
	}
	v.setupCustomValidation()
	return v
}

func (v *validatorClient) setupCustomValidation() {
	v.validator.RegisterValidation("win-lose", checkIsTeamResultValid)
	v.validator.RegisterValidation("parseint", checkIsPublishedAtValid)
	v.validator.RegisterValidation("check-limit", checkIsLimitValid)
}

func (v *validatorClient) Validate(s interface{}) error {
	err := v.validator.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return ErrInvalidation
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}
		return ErrInvalidation
	}
	return nil
}
