package router

import (
	"net/http"

	"redler/docs"
	"redler/internal/middlewares"
	"redler/internal/modules/user"
	"redler/internal/pkg/config"
	"redler/internal/pkg/lib"
	"redler/internal/pkg/version"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(engine *gin.Engine, config *config.Config, db *gorm.DB) {
	rg := engine.Group(config.BasePath)

	addSwaggerEndpoint(rg, config)

	rg.Use(middlewares.JSONLogMiddleware())
	rg.GET("/version", versionHandlerShort)

	user.InitRoutes("/user", rg, config, db)
}

func versionHandlerShort(c *gin.Context) {
	versionData := version.Short()
	lib.HttpResponse(http.StatusOK).Data(versionData).Send(c)
}

func addSwaggerEndpoint(rg *gin.RouterGroup, config *config.Config) {
	docs.SwaggerInfo.Version = version.GetVersion()
	docs.SwaggerInfo.BasePath = config.BasePath
	rg.GET("/apidocs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
