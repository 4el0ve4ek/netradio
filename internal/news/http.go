package news

import (
	"fmt"
	"net/http"
	"netradio/libs/jwt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"netradio/internal/databases/news"
	"netradio/pkg/handles"
	"netradio/pkg/log"
)

func NewHTTPServant(config Config, logger log.Logger, newsService news.Service, verificator jwt.Verificator) *servant {
	server := &http.Server{}
	server.Addr = fmt.Sprintf(":%d", config.Port)

	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	core := handles.NewCore(logger, verificator)

	addJSONHandler(core, router, "GET", "/news", newRecentGetterHandler(newsService))
	addJSONHandler(core, router, "GET", "/news/{newsID}", newGetterHandler(newsService))

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
