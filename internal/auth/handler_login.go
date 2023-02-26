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

func (h *loginHandler) ServeHTTP(context context.Context, request *http.Request) (handles.Response, error) {
	var rawUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&rawUser)
	if err != nil {
		return handles.Response{}, err
	}

	user := h.userService.GetUser(rawUser.Email, rawUser.Password)
	resp := handles.Response{
		Headers: make(http.Header),
	}

	err = h.verificator.AddUIDToHeader(resp.Headers, user)
	return resp, err
}
