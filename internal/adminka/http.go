package adminka

import (
	"netradio/internal/databases/music"
	"netradio/internal/databases/news"
	"netradio/pkg/handles"

	"github.com/go-chi/chi/v5"
)

///go:embed static
//var content embed.FS

func RoutePaths(
	core handles.Core,
	router chi.Router,
	newsService news.Service,
	musicService music.Service,
) {

	addHandler(core, router, "POST", "/news/add", newCreateNewsHandler(newsService))
	addHandler(core, router, "POST", "/news/change", newChangeNewsHandler(newsService))
	addHandler(core, router, "POST", "/podcasts/add", newCreatePodcastHandler(musicService))
	addHandler(core, router, "POST", "/podcasts/change", newChangePodcastHandler(musicService))

	//fsys, _ := fs.Sub(content, "static")
	//router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(fsys))))
}

func addHandler(core handles.Core, router chi.Router, method, pattern string, handler handles.Handler) {
	router.Method(method, pattern, handles.NewAuthWrapper(handles.NewResponseWrapper(handler), core))
}
