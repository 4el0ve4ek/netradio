package context

import (
	"net/http"
	"netradio/internal/databases/user"
	"netradio/libs/jwt"
	"netradio/models"
	"netradio/pkg/log"
)

type Context interface {
	GetUser() models.User
	GetLogger() log.Logger
	GetRequest() *http.Request
}

func New(r *http.Request, verificator jwt.Verificator, userService user.Service, logger log.Logger) *context {
	return &context{
		request:     r,
		verificator: verificator,
		userService: userService,
		logger:      logger,
	}
}

type context struct {
	request     *http.Request
	verificator jwt.Verificator
	userService user.Service
	logger      log.Logger
}

func (c *context) GetUser() models.User {
	uid, err := c.verificator.GetUIDFromHeader(c.request.Header)
	if err != nil {
		c.logger.Warn(err)
	}
	user, err := c.userService.GetUserByUID(uid)
	if err != nil {
		c.logger.Warn(err)
	}

	return user
}

func (c *context) GetLogger() log.Logger {
	return c.logger
}

func (c *context) GetRequest() *http.Request {
	return c.request
}
