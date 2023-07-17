package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/bohoslavskyi/ip-info/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	httpServer := new(server.Server)
	go func() {
		if err := httpServer.Run(8000); err != nil {
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
