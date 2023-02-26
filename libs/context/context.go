package context

import (
	"net/http"
	"netradio/libs/jwt"
	"netradio/models"
	"netradio/pkg/log"
)

type Context interface {
	GetUser() models.User
	GetLogger() log.Logger
}

func New(r *http.Request, verificator jwt.Verificator, logger log.Logger) *context {
	return &context{
		request:     r,
		verificator: verificator,
		logger:      logger,
	}
}

type context struct {
	request     *http.Request
	verificator jwt.Verificator
	logger      log.Logger
}

func (c *context) GetUser() models.User {
	uid, err := c.verificator.GetUIDFromHeader(c.request.Header)
	if err != nil {
		c.logger.Warn(err)
	}

	return models.User{
		UID: uid,
		//Nickname: "Ivan",
		//Status:   models.UserRegistered,
	}
}

func (c *context) GetLogger() log.Logger {
	return c.logger
}
