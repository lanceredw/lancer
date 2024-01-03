package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(r *gin.Engine) {

	//swag api
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//self api
	api := r.Group("/api")
	ApiRouter(api)
}
