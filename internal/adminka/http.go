package adminka

import (
	"fmt"
	"net/http"
	"netradio/internal/databases/news"
	"netradio/libs/jwt"

	"netradio/pkg/handles"
	"netradio/pkg/log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHTTPServant(config Config, logger log.Logger, verificator jwt.Verificator, newsService news.Service) *servant {
	server := &http.Server{}
	server.Addr = fmt.Sprintf(":%d", config.Port)

	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)

	core := handles.NewCore(logger, verificator)

	addHandler(core, router, "POST", "/news/add", newCreateHandler(newsService))
	addHandler(core, router, "POST", "/news/change", newModifyHandler(newsService))

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
