package podcast

import (
	"netradio/internal/databases/music"
	"netradio/pkg/handles"

	"github.com/go-chi/chi/v5"
)

func RoutePaths(
	core handles.Core,
	router chi.Router,
	musicService music.Service,
) {
	addJSONHandler(core, router, "GET", "/podcast/all", newGetterHandler(musicService))
	//addStreamingHandler(core, router, "GET", "/podcast/{podcastID}/start", newInfoGetterHandler(channelService))
	//addJSONHandler(core, router, "GET", "/broadcast/{broadcastID}/program", newProgramGetterHandler(channelService))
	//addStreamingHandler(core, router, "GET", "/broadcast/{broadcastID}/start", newStartHandler(channelService)) // <- streaming handler

}

func addJSONHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.HandlerJSON) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewResponseWrapper(handles.NewJSONWrapper(handler)), core))
}

func addHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.Handler) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewResponseWrapper(handler), core))
}
