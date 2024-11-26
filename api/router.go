package api

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"ryg-api-gateway/api/handlers"
	"ryg-api-gateway/api/token"
	_ "ryg-api-gateway/docs"
)

// @tag.name Authentication
// @tag.description User registration and login methods

// @tag.name User
// @tag.description User methods

// @tag.name Challenge
// @tag.description Challenge methods

// @tag.name Task
// @tag.description Task methods

// NewGinRouter creates a new gin router with the specified RpcClientManager
// @Title RYG API Gateway
// @Version 1.0
// @Description This is the API Gateway for the RYG project.
// @SecurityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func NewGinRouter(cm *handlers.RpcClientManager) *gin.Engine {
	router := gin.Default()

	swaggerUrl := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, swaggerUrl))

	// Auth routes
	router.POST("/login", cm.Login)
	router.POST("/register", cm.RegisterUser)

	// User routes
	userGroup := router.Group("/users")
	userGroup.Use(token.JWTMiddleware())
	userGroup.GET("", cm.GetProfile)
	userGroup.PUT("", cm.UpdateUser)
	userGroup.DELETE("", cm.DeleteUser)

	// Challenge routes
	challengeGroup := router.Group("/challenges")
	challengeGroup.Use(token.JWTMiddleware())

	challengeGroup.POST("/", cm.CreateChallenge)
	challengeGroup.GET("/:challenge_id", cm.GetChallenge)
	challengeGroup.GET("/", cm.GetChallenges)
	challengeGroup.PUT("/:challenge_id", cm.UpdateChallenge)
	challengeGroup.DELETE("/:challenge_id", cm.DeleteChallenge)
	challengeGroup.POST("/:challenge_id/start", cm.StartChallenge)
	challengeGroup.POST("/:challenge_id/finish", cm.FinishChallenge)
	challengeGroup.POST("/:challenge_id/invite", cm.InviteUser)
	challengeGroup.DELETE("/:challenge_id/unsubscribe", cm.UnsubscribeFromChallenge)
	router.GET("/challenges/accept", cm.AcceptChallengeInvitation)

	// Task routes
	taskGroup := challengeGroup.Group("/:challenge_id/tasks")
	taskGroup.POST("/", cm.CreateTask)
	taskGroup.POST("/bulk", cm.CreateTasks)
	taskGroup.GET("/", cm.GetTasks)
	taskGroup.GET("/:task_id", cm.GetTask)
	taskGroup.PUT("/:task_id", cm.UpdateTask)
	taskGroup.DELETE("/:task_id", cm.DeleteTask)
	taskGroup.PUT("/:task_id/status", cm.UpdateTaskStatus)

	return router
}
