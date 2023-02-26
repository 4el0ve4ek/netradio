package auth

import (
	"fmt"
	"net/http"
	"netradio/internal/databases/user"
	"netradio/libs/jwt"
	"netradio/pkg/handles"
	"netradio/pkg/log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHTTPServant(config Config, logger log.Logger, verificator jwt.Verificator, userService user.Service) *servant {
	server := &http.Server{}
	server.Addr = fmt.Sprintf(":%d", config.Port)

	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)

	core := handles.NewCore(logger, verificator)

	addHandler(core, router, "POST", "/register", newRegisterHandler(userService, verificator))
	addHandler(core, router, "POST", "/login", newLoginHandler(userService, verificator))
	addHandler(core, router, "GET", "/refresh", newRefreshHandler(verificator))

	addJSONHandler(core, router, "GET", "/user", newUserGetterHandler(userService))
	addHandler(core, router, "POST", "/user", newUserChangeHandler(userService))
	addHandler(core, router, "DELETE", "/user", newUserDeleteHandler(userService, verificator))

	server.Handler = router
	return &servant{
		server: server,
	}
}

type servant struct {
	server *http.Server
}

func (s *servant) GetServer() *http.Server {
	return s.server
}

func addHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.Handler) {
	router.Method(method, pattern, handles.NewAuthWrapper(handler, core))
}

func addJSONHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.HandlerJSON) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewJSONWrapper(handler), core))
}
