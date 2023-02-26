package podcast

import (
	"fmt"
	"net/http"
	"netradio/internal/databases/music"
	"netradio/pkg/log"

	"netradio/libs/jwt"
	"netradio/pkg/handles"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHTTPServant(
	config Config,
	logger log.Logger,
	verificator jwt.Verificator,
	musicService music.Service,
) *servant {
	server := &http.Server{}
	server.Addr = fmt.Sprintf(":%d", config.Port)

	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	core := handles.NewCore(logger, verificator)
	_ = core
	//addJSONHandler(core, router, "GET", "/podcast/all", nil)
	//addStreamingHandler(core, router, "GET", "/podcast/{podcastID}/start", newInfoGetterHandler(channelService))
	//addJSONHandler(core, router, "GET", "/broadcast/{broadcastID}/program", newProgramGetterHandler(channelService))
	//addSocketHandler(core, router, "GET", "/broadcast/{broadcastID}/start", newStartHandler(channelService)) // <- streaming handler

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

func addJSONHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.HandlerJSON) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewJSONWrapper(handler), core))
}

func addHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.Handler) {
	router.Method(method, pattern, handles.NewAuthWrapper(handler, core))
}
