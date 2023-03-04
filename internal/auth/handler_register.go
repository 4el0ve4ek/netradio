package auth

import (
	"encoding/json"
	"net/http"
	"netradio/internal/databases/user"
	"netradio/libs/context"
	"netradio/libs/jwt"
	"netradio/pkg/handles"
)

func newRegisterHandler(userService user.Service, authService jwt.Verificator) *registerHandler {
	return &registerHandler{
		userService: userService,
		authService: authService,
	}
}

type registerHandler struct {
	userService user.Service
	authService jwt.Verificator
}

func (h *registerHandler) ServeHTTP(context context.Context) (handles.Response, error) {
	request := context.GetRequest()
	var rawUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&rawUser)
	if err != nil {
		return handles.Response{}, err
	}

	user, err := h.userService.AddUser(rawUser.Email, rawUser.Password)
	resp := handles.Response{
		Headers: make(http.Header),
	}

	return resp, h.authService.AddUIDToHeader(resp.Headers, user)
}
