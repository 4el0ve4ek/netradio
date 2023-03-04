package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"netradio/cmd/config"
	"netradio/internal/adminka"
	"netradio/internal/databases/music"
	newsdb "netradio/internal/databases/news"
	"netradio/internal/databases/user"
	"netradio/libs/jwt"
	"netradio/pkg/handles"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

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

	newsDB := newsdb.NewService()
	userDB := user.NewService()
	musicDB := music.NewService()

	router := chi.NewRouter()

	core := handles.NewCore(logger, verificator, userDB)
	adminka.RoutePaths(core, router, newsDB, musicDB)

	server := &http.Server{}
	server.Addr = fmt.Sprintf(":%d", cfg.Adminka.Port)
	server.Handler = router

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(errors.Wrap(err, "http server failure"))
			sigChan <- syscall.SIGINT
		}
	}()

	<-sigChan
	logger.Info("shutting down")

}
