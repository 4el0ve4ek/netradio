package auth

import (
	"netradio/internal/databases/user"
	"netradio/libs/jwt"
	"netradio/pkg/handles"

	"github.com/go-chi/chi/v5"
)

func RoutePaths(
	core handles.Core,
	router chi.Router,
	verificator jwt.Verificator,
	userService user.Service,
) {

	addHandler(core, router, "POST", "/register", newRegisterHandler(userService, verificator))
	addHandler(core, router, "POST", "/login", newLoginHandler(userService, verificator))
	addHandler(core, router, "GET", "/refresh", newRefreshHandler(verificator))

	addJSONHandler(core, router, "GET", "/user", newUserGetterHandler(userService))
	addHandler(core, router, "POST", "/user", newUserChangeHandler(userService))
	addHandler(core, router, "DELETE", "/user", newUserDeleteHandler(userService, verificator))
}

func addHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.Handler) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewResponseWrapper(handler), core))
}

func addJSONHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.HandlerJSON) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewResponseWrapper(handles.NewJSONWrapper(handler)), core))
}
