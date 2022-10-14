package server

import (
	"github.com/gin-gonic/gin"
	"github.com/prakharporwal/go-server/pkg/klog"
)

type Server struct {
	router *gin.Engine
}

func NewServer(initRouter func() *gin.Engine) *Server {
	server := &Server{
		router: initRouter(),
	}
	return server
}

func (server *Server) Start(serverAddress string, serverPort string) error {
	klog.Info("Starting Server!")
	return server.router.Run(serverAddress + ":" + serverPort)
}

func (server *Server) Stop() {
	klog.Info("Stopping Server!")
}
