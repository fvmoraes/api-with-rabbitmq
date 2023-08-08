package helpers

import (
	"github.com/fvmoraes/api-with-rabbitmq/internal/models"

	"gopkg.in/validator.v2"
)

func ModelValidator(foobar *models.Foobar) error {
	if err := validator.Validate(foobar); err != nil {
		return err
	}
	return nil
}
