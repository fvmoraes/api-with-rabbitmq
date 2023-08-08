package routes

import (
	"github.com/fvmoraes/api-with-rabbitmq/internal/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func MyFoobarRoutes(r *gin.Engine) {
	g := r.Group("/api/v1")
	g.GET("/foobar", controllers.ShownFoobar)
	g.GET("/foobar/:id", controllers.ShownFoobarByParamId)
	g.POST("/foobar", controllers.CreateFoobar)
	g.PATCH("/foobar/:id", controllers.EditFoobarByParamId)
	g.DELETE("/foobar/:id", controllers.DeleteFoobarByParamId)
	g.GET("/foobar/param/:reg", controllers.ShownFoobarByParamReg)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
