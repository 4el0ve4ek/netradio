package main

import (
	"math/rand"
	stdhttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"netradio/cmd/config"
	"netradio/internal/broadcast"
	"netradio/internal/databases/channel"
	"netradio/libs/jwt"
	"netradio/pkg/errors"
	"netradio/pkg/log"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	logger := log.NewLogger()

	cfg, err := config.ReadYAML(config.DefaultYAMLPath)
	if err != nil {
		logger.Fatal(err)
	}
	verificator := jwt.NewVerificator(cfg.Jwt)
	broadcastDB := channel.NewService()
	servant := broadcast.NewHTTPServant(cfg.Broadcast, logger, verificator, broadcastDB)

	server := servant.GetServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != stdhttp.ErrServerClosed {
			logger.Error(errors.Wrap(err, "http server failure"))
			sigChan <- syscall.SIGINT
		}
	}()

	<-sigChan
	logger.Info("shutting down")

}
