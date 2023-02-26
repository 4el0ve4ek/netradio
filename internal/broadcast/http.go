package broadcast

import (
	"fmt"
	"net/http"
	"netradio/internal/databases/channel"
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
	channelService channel.Service,
) *servant {
	server := &http.Server{}
	server.Addr = fmt.Sprintf(":%d", config.Port)

	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	core := handles.NewCore(logger, verificator)

	addJSONHandler(core, router, "GET", "/broadcast/all", newGetterHandler(channelService))
	addJSONHandler(core, router, "GET", "/broadcast/{broadcastID}/info", newInfoGetterHandler(channelService))
	//addJSONHandler(core, router, "GET", "/broadcast/{broadcastID}/program", newProgramGetterHandler(channelService))
	//addSocketHandler(core, router, "GET", "/broadcast/{broadcastID}/start", newStartHandler(channelService))

	// moved to adminka
	//addHandler(core, router, "POST", "/news/add", adminka.newCreateHandler(newsService))
	//addHandler(core, router, "POST", "/news/modify", adminka.newModifyHandler(newsService))

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
