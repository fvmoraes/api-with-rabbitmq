package main

import (
	"github.com/fvmoraes/api-with-rabbitmq/internal/initializers"
	"github.com/fvmoraes/api-with-rabbitmq/internal/servers"
)

func main() {
	initializers.StartDatabaseConnect()
	initializers.StartMigration()
	servers.RunMyFoobarServer()
}
