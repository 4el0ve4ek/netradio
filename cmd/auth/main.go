package main

import (
	"fmt"
	"math/rand"
	"net/http"
	stdhttp "net/http"
	"netradio/libs/jwt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

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
	userService := user.NewService()

	router := chi.NewRouter()
	auth.RoutePaths(cfg.Auth, router, logger, verificator, userService)

	server := &http.Server{}
	server.Addr = fmt.Sprintf(":%d", cfg.Auth.Port)
	server.Handler = router

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
