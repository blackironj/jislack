package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackironj/jislack/config"
	"github.com/blackironj/jislack/slacktool"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitCfg("config/config.yaml")

	router := gin.Default()

	router.Use(slacktool.ValidateSlackCommandMiddleware())
	router.POST("/jislack", slacktool.CommandHandler)

	cfg := config.Get()
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown: ", err)
	}
	log.Println("Server exiting")
}
