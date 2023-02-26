package handles

import (
	"netradio/libs/jwt"
	"netradio/pkg/log"
)

type Core struct {
	auth jwt.Verificator
	log  log.Logger
}

func NewCore(log log.Logger, auth jwt.Verificator) Core {
	return Core{
		auth: auth,
		log:  log,
	}
}
