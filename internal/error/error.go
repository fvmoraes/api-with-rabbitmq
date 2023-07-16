package error

import (
	"log"

	"github.com/fvmoraes/api-with-rabbitmq/internal/logs"
)

func ValidateError(msg string, err error) {
	if err != nil {
		logs.WriteLogFile("ERROR", msg+" "+err.Error())
		log.Fatalf(msg, err)
	}
}
