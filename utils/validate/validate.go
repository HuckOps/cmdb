package validate

import "github.com/go-playground/validator/v10"

func Validate(data interface{}) error  {
	v := validator.New(validator.WithRequiredStructEnabled())
	return v.Struct(data)
}
