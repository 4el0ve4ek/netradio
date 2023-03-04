package news

import (
	"netradio/internal/databases/news"
	"netradio/pkg/handles"

	"github.com/go-chi/chi/v5"
)

func RoutePaths(
	core handles.Core,
	router chi.Router,
	newsService news.Service,
) {
	addJSONHandler(core, router, "GET", "/news", newRecentGetterHandler(newsService))
	addJSONHandler(core, router, "GET", "/news/{newsID}", newGetterHandler(newsService))
}

func addJSONHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.HandlerJSON) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewResponseWrapper(handles.NewJSONWrapper(handler)), core))
}

func addHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.Handler) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewResponseWrapper(handler), core))
}
