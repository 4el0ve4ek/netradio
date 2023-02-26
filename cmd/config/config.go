package config

import (
	"netradio/internal/adminka"
	"netradio/internal/auth"
	"netradio/internal/broadcast"
	"netradio/internal/news"
	"netradio/internal/podcast"
	"netradio/libs/jwt"
)

type Config struct {
	Adminka   adminka.Config   `yaml:"adminka"`
	Auth      auth.Config      `yaml:"auth"`
	Broadcast broadcast.Config `yaml:"broadcast"`
	Jwt       jwt.Config       `yaml:"jwt"`
	News      news.Config      `yaml:"news"`
	Podcast   podcast.Config   `yaml:"podcast"`
}
