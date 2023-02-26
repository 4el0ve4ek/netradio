package auth

import (
	"net/http"
	"netradio/libs/context"
	"netradio/libs/jwt"
	"netradio/pkg/handles"
)

func newRefreshHandler(authService jwt.Verificator) *refreshHandler {
	return &refreshHandler{authService: authService}
}

type refreshHandler struct {
	authService jwt.Verificator
}

func (h *refreshHandler) ServeHTTP(context context.Context, request *http.Request) (handles.Response, error) {
	user := context.GetUser()
	resp := handles.Response{
		Headers: make(http.Header),
	}

	if user.UID == 0 {
		h.authService.DeleteAuth(resp.Headers)
	} else {
		err := h.authService.AddUIDToHeader(resp.Headers, user)
		return resp, err
	}

	return resp, nil
}
