package api

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"ryg-api-gateway/api/handlers"
	_ "ryg-api-gateway/docs"
)

// NewGinRouter
// @title RYG API Gateway
// @version 1.0
// @description This is the API Gateway for the RYG project.
func NewGinRouter(cm *handlers.RpcClientManager) *gin.Engine {
	router := gin.Default()

	swaggerUrl := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, swaggerUrl))

	userGroup := router.Group("/users")
	userGroup.POST("/register", cm.RegisterUser)
	userGroup.GET("/:id", cm.GetProfile)
	userGroup.PUT("/", cm.UpdateUser)
	userGroup.DELETE("/:id", cm.DeleteUser)

	return router
}
