package auth

import (
	"encoding/json"
	"net/http"

	"netradio/internal/databases/user"
	"netradio/libs/context"
	"netradio/models"
	"netradio/pkg/handles"
)

func newUserChangeHandler(userService user.Service) *userChangeHandler {
	return &userChangeHandler{userService: userService}
}

type userChangeHandler struct {
	userService user.Service
}

func (h *userChangeHandler) ServeHTTP(context context.Context, request *http.Request) (handles.Response, error) {
	var rawUser struct {
		Nickname  *string      `json:"name,omitempty"`
		PhotoLink *string      `json:"photo,omitempty"`
		Lang      *models.Lang `json:"lang,omitempty"`
		Password  *string      `json:"password,omitempty"`
	}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&rawUser)
	if err != nil {
		return handles.Response{}, err
	}
	return handles.Response{}, nil
}
