package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"netradio/pkg/handles"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"

	"netradio/cmd/config"
	"netradio/internal/adminka"
	"netradio/internal/auth"
	"netradio/internal/broadcast"
	"netradio/internal/databases/channel"
	"netradio/internal/databases/music"
	newsdb "netradio/internal/databases/news"
	"netradio/internal/databases/user"
	"netradio/internal/news"
	"netradio/internal/podcast"
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

	newsDB := newsdb.NewService()
	musicDB := music.NewService()
	broadcastDB := channel.NewService()
	userDB := user.NewService()

	router := chi.NewRouter()

	core := handles.NewCore(logger, verificator, userDB)
	adminka.RoutePaths(core, router, newsDB, musicDB)
	auth.RoutePaths(core, router, verificator, userDB)
	broadcast.RoutePaths(core, router, broadcastDB)
	news.RoutePaths(core, router, newsDB)
	podcast.RoutePaths(core, router, musicDB)

	server := &http.Server{}
	server.Addr = fmt.Sprintf(":%d", 80)
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
