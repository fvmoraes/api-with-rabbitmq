package initializers

import "github.com/fvmoraes/api-with-rabbitmq/internal/models"

func StartMigration() {
	DB.AutoMigrate(&models.Foobar{})
}
