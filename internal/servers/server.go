package servers

import (
	"github.com/fvmoraes/api-with-rabbitmq/internal/docs"
	"github.com/fvmoraes/api-with-rabbitmq/internal/routes"

	"github.com/gin-gonic/gin"
)

func RunMyFoobarServer() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	routes.MyFoobarRoutes(r)
	r.Run(":9000")
}
