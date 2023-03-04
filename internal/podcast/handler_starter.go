package podcast

import (
	"encoding/json"
	"net/http"
	"netradio/internal/databases/music"
	"netradio/libs/context"
)

func newStreamingHandler(musicService music.Service) *streamingHandler {
	return &streamingHandler{
		musicService: musicService,
	}
}

type streamingHandler struct {
	musicService music.Service
}

func (h *getterHandler) ServeHTTP(context context.Context, request *http.Request) (json.RawMessage, error) {
	podcasts := h.musicService.GetPodcasts()

	return json.Marshal(podcasts)
}
