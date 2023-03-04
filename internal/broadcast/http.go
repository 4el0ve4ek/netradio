package broadcast

import (
	"net/http"
	"netradio/internal/databases/channel"
	"netradio/pkg/handles"

	"github.com/go-chi/chi/v5"
)

func RoutePaths(
	core handles.Core,
	router chi.Router,
	channelService channel.Service,
) {

	addJSONHandler(core, router, "GET", "/broadcast/all", newGetterHandler(channelService))
	addJSONHandler(core, router, "GET", "/broadcast/{broadcastID}/info", newInfoGetterHandler(channelService))
	//addJSONHandler(core, router, "GET", "/broadcast/{broadcastID}/program", newProgramGetterHandler(channelService))
	//addSocketHandler(core, router, "GET", "/broadcast/{broadcastID}/start", newStartHandler(channelService))
}

type servant struct {
	server *http.Server
}

func (s *servant) GetServer() *http.Server {
	return s.server
}

func addJSONHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.HandlerJSON) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewResponseWrapper(handles.NewJSONWrapper(handler)), core))
}

func addHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.Handler) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewResponseWrapper(handler), core))
}
