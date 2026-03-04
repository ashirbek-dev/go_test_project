package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	//"incard.uz/humo/http/controllers"
)

func (s *Server) Routes() *gin.Engine {
	_router := s.router

	//router.Use(middlewares.Locale())

	_router.Use(cors.Default())

	controller := Controller{
		appService: s.appService,
	}

	api := _router.Group("/api")
	//api.Use(BasicAuthMiddleware())

	api.POST("v1", controller.PostHandler)
	api.GET("v1", controller.GetHandler)

	return _router
}
