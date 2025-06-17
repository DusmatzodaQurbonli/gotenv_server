package routes

import (
	_ "Gotenv/docs"
	"Gotenv/internal/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// InitRoutes — настраиваем HTTP-маршруты
func InitRoutes(r *gin.Engine) *gin.Engine {
	// Health-check
	r.GET("/ping", controllers.Ping)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
