package routes

import (
	_ "Gotenv/docs"
	"Gotenv/internal/controllers"
	"Gotenv/internal/controllers/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRoutes — настраиваем HTTP-маршруты
func InitRoutes(r *gin.Engine) *gin.Engine {
	// Health-check
	r.GET("/ping", controllers.Ping)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth
	auth := r.Group("/auth")
	{
		auth.POST("/sign-up", controllers.SignUp)
		auth.POST("/sign-in", controllers.SignIn)
		auth.POST("/refresh", controllers.RefreshToken)

		auth.GET("/project", controllers.GetIPProject)
	}

	project := r.Group("/projects", middlewares.CheckUserAuthentication)
	{
		project.GET("", controllers.GetProjectsUser)
		project.GET("/:id", controllers.GetProjectByIDAndUserID)
		project.POST("", controllers.CreateProject)
		project.PUT("/:id", middlewares.CheckUsersProject, controllers.UpdateProject)
		project.PATCH("/:id/active", middlewares.CheckUsersProject, controllers.UpdateProjectActive)
		project.DELETE("/:id", middlewares.CheckUsersProject, controllers.DeleteProject)
	}

	projectVars := project.Group("/vars")
	{
		projectVars.POST("/val/:id", middlewares.CheckUsersProject, controllers.GetAllProjectVars)
		projectVars.POST("/:id", middlewares.CheckUsersProject, controllers.CreateProjectVars)
		projectVars.PUT("/:id", middlewares.CheckUsersProject, controllers.UpdateProjectVars)
		projectVars.DELETE("/:id", middlewares.CheckUsersProjectByVarsID, controllers.DeleteProjectVars)
	}

	return r
}
