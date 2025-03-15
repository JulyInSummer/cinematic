package validators

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func DateFormat(fl validator.FieldLevel) bool {
	const format = "2006"

	value := fl.Field().String()
	_, err := time.Parse(format, value)

	return err == nil
}
