package auth

import (
	"encoding/json"
	"net/http"
	"netradio/internal/databases/user"
	"netradio/libs/context"
	"netradio/libs/jwt"
	"netradio/pkg/handles"
)

func newLoginHandler(userService user.Service, authService jwt.Verificator) *loginHandler {
	return &loginHandler{
		userService: userService,
		verificator: authService,
	}
}

type loginHandler struct {
	userService user.Service
	verificator jwt.Verificator
}

func (h *loginHandler) ServeHTTP(context context.Context) (handles.Response, error) {
	var rawUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(context.GetRequest().Body)
	err := decoder.Decode(&rawUser)
	if err != nil {
		return handles.Response{}, err
	}

	user, err := h.userService.GetUser(rawUser.Email, rawUser.Password)
	resp := handles.Response{
		Headers: make(http.Header),
	}

	if err != nil {
		return resp, err
	}
	err = h.verificator.AddUIDToHeader(resp.Headers, user)
	return resp, err
}
