package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// Server is a struct which holds a http server state
type Server struct {
	srv *http.Server
}

// Create will take in a gin engine and will return a server struct ready to be run
func Create(router *gin.Engine) *Server {

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	return &Server{srv: srv}

}

func (srv *Server) Run() {

	go func() {
		if err := srv.srv.ListenAndServe(); err != nil {
			log.Printf("server running")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown")
	}

	log.Println("server exiting")
}
