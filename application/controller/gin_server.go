package controller

import (
	"bee-pod-master/pkg/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type ginServer struct {
	router *gin.Engine
}

type GinServer interface {
	SetupRouter(es MasterController) *gin.Engine
}

func NewGinServer() GinServer {
	return &ginServer{
		router: gin.New(),
	}
}

func (s *ginServer) SetupRouter(es MasterController) *gin.Engine {
	s.router.Use(gzip.Gzip(gzip.BestCompression))
	s.router.Use(middleware.GlobalRecover())
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.router.GET("", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
		c.Abort()
	})

	s.router.POST("/master/copy", es.Copy())
	s.router.GET("/master/getenv", es.GetEnv())
	s.router.GET("/master/getconfig", es.GetConfig())
	s.router.POST("/master/search", es.Search())
	s.router.POST("/master/partition", es.Partition())
	s.router.DELETE("/master", es.Delete())
	s.router.GET("/_monitoring/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, "OK")
	})

	return s.router
}
