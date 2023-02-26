package auth

import (
	"encoding/json"
	"net/http"
	"netradio/internal/databases/user"
	"netradio/libs/context"
)

func newUserGetterHandler(userService user.Service) *userGetterHandler {
	return &userGetterHandler{
		userService: userService,
	}
}

type userGetterHandler struct {
	userService user.Service
}

func (h *userGetterHandler) ServeJSON(context context.Context, request *http.Request) (json.RawMessage, error) {
	return json.Marshal(context.GetUser())
}
