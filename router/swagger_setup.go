package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Possible routes for accessing swagger are:
// '/docs', '/swagger', '/docs/*', '/swagger/index.html'
func setupSwagger(e *gin.Engine) {
	e.GET("/swagger", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
	e.GET("/docs/*any", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
	})
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
