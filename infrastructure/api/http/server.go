package http

import (
	"fmt"
	"gateway/core/app"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	appService *app.ApplicationService
	router     *gin.Engine
}

func CreateServer(appService *app.ApplicationService) *Server {
	gin.SetMode(gin.ReleaseMode)
	/*_router := gin.Default()
	_router.Use(cors.Default())*/
	_router := gin.New()
	_router.Use(gin.Recovery())
	return &Server{
		appService: appService,
		router:     _router,
	}
}

func (s *Server) Run(port int) error {
	r := s.Routes()
	err := r.Run(fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}
	return nil
}
