package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bohoslavskyi/ip-info/configs"
	"github.com/bohoslavskyi/ip-info/internal/handler"
	"github.com/bohoslavskyi/ip-info/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	cfg, err := configs.NewConfig()
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	handlers := handler.NewHandler()
	httpServer := new(server.Server)
	go func() {
		if err := httpServer.Run(cfg.ServerPort, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running HTTP server: %s.", err.Error())
		}
	}()
	logrus.Print("HTTP server has been started.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("HTTP server is shutting down.")
	if err := httpServer.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s.\n", err.Error())
	}
}
