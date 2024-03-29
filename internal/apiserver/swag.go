package apiserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"goer-startup/api/swagger/docs"
)

// MapSwagRoutes
//
//	@title						API Docs
//	@version					1.0
//	@BasePath					/
//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
func MapSwagRoutes(r *gin.Engine) {
	// swagger info
	docs.SwaggerInfo.Title = "API Docs"
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// use ginSwagger middleware to serve the API docs
	// DefaultModelsExpandDepth: set -1 to hide models below
	r.GET("/api/docs/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))
}
