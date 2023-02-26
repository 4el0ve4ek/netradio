package main

import (
	"math/rand"
	stdhttp "net/http"
	"netradio/libs/jwt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"netradio/cmd/config"
	"netradio/internal/auth"
	"netradio/internal/databases/user"
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
	userService := user.NewService(logger)

	servant := auth.NewHTTPServant(cfg.Auth, logger, verificator, userService)

	server := servant.GetServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	logger.Info("starts on port " + strconv.Itoa(cfg.Auth.Port))
	go func() {
		if err := server.ListenAndServe(); err != nil && err != stdhttp.ErrServerClosed {
			logger.Error(errors.Wrap(err, "http server failure"))
			sigChan <- syscall.SIGINT
		}
	}()

	<-sigChan
	logger.Info("shutting down")

}
