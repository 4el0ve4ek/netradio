package handles

import (
	"netradio/internal/databases/user"
	"netradio/libs/jwt"
	"netradio/pkg/log"
)

type Core struct {
	auth        jwt.Verificator
	userService user.Service
	log         log.Logger
}

func NewCore(log log.Logger, auth jwt.Verificator, userService user.Service) Core {
	return Core{
		auth:        auth,
		log:         log,
		userService: userService,
	}
}
