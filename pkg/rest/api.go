package rest

import (
	"danu/danu-server/pkg/config"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Server struct {
	config *config.Config
	engine *gin.Engine
}

func NewServer(cfg *config.Config, e *gin.Engine) *Server {
	rand.Seed(time.Now().UTC().UnixNano())

	return &Server{
		config: cfg,
		engine: e,
	}
}

func (s *Server) Initialise() {

	s.engine.POST("api/placeholder", s.PlaceholderFunc)

	var filename = "logfile.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	log.SetFormatter(&log.JSONFormatter{})
	if err != nil {
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}

}

func (s *Server) PlaceholderFunc(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}
