package auth

import (
	"net/http"
	"netradio/internal/databases/user"
	"netradio/libs/context"
	"netradio/libs/jwt"
	"netradio/pkg/handles"
)

func newUserDeleteHandler(userService user.Service, verificator jwt.Verificator) *userDeleteHandler {
	return &userDeleteHandler{
		userService: userService,
		verificator: verificator,
	}
}

type userDeleteHandler struct {
	userService user.Service
	verificator jwt.Verificator
}

func (h *userDeleteHandler) ServeHTTP(context context.Context) (handles.Response, error) {
	h.userService.Delete(context.GetUser())

	resp := handles.Response{
		Headers: make(http.Header),
	}
	h.verificator.DeleteAuth(resp.Headers)
	return resp, nil
}
