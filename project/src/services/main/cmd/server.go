package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"echo-inertia.com/src/services/main/internal/bootstrap"
	"echo-inertia.com/src/services/main/internal/handler/router"
)

type Server struct {
	Bootstrap *bootstrap.Bootstrap
	Srv       *http.Server
	Addr      string
}

func NewServer() *Server {
	b := bootstrap.MustGetBootstrapInstance()
	addr := fmt.Sprintf(":%s", b.Config.Port)
	router := router.NewRouter(b)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	return &Server{
		Bootstrap: b,
		Srv:       srv,
		Addr:      addr,
	}
}

func Run() {
	runHTTPServer(NewServer())
}

func runHTTPServer(s *Server) {
	env := "DEVELOPMENT"
	delay := 500 * time.Millisecond
	if s.Bootstrap.Config.Env != "development" {
		env = "PRODUCTION"
		delay = 3 * time.Second
	}
	s.Bootstrap.RunJobs()

	log.Printf("HTTP server running at :%s", s.Bootstrap.Config.Port)
	log.Printf("( %s MODE ) running main project", env)
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGINT,
	)
	defer stop()
	go func() {
		if err := s.Srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP Server Error %v", err)
		}
	}()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), delay)
	defer cancel()
	if err := s.Srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown error: %v\n", err)
	}
}

func RunLoop() {
	for {
		log.Println("Running server for ever")
		time.Sleep(1 * time.Minute)
	}
}
