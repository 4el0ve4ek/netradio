package podcast

import (
	"encoding/json"
	"netradio/internal/databases/music"
	"netradio/libs/context"
)

func newGetterHandler(musicService music.Service) *getterHandler {
	return &getterHandler{
		musicService: musicService,
	}
}

type getterHandler struct {
	musicService music.Service
}

func (h *getterHandler) ServeJSON(context context.Context) (json.RawMessage, error) {
	podcasts := h.musicService.GetPodcasts()

	return json.Marshal(podcasts)
}
