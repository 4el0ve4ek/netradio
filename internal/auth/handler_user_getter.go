package auth

import (
	"encoding/json"
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

func (h *userGetterHandler) ServeJSON(context context.Context) (json.RawMessage, error) {
	return json.Marshal(context.GetUser())
}
