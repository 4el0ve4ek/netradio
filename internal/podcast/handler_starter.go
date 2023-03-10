package podcast

import (
	"fmt"
	"net/http"
	"netradio/internal/databases/music"
	"netradio/libs/context"
	"netradio/models"
	"netradio/pkg/errors"
)

func newPodcastGetter(musicService music.Service) *podcastGetterHandler {
	return &podcastGetterHandler{
		musicService: musicService,
	}
}

type podcastGetterHandler struct {
	musicService music.Service
}

func (h *podcastGetterHandler) ServeHTTP(context context.Context, w http.ResponseWriter) error {

	musicChunks, err := music.NewService().LoadMusicBatch(models.MusicInfo{})
	if err != nil {
		return errors.Wrap(err, "load music batch")
	}

	f, ok := w.(http.Flusher)
	if !ok {
		return errors.New("connection not flushable")
	}

	headers := w.Header()
	headers.Set("Content-Type", "text/event-stream")
	headers.Set("Cache-Control", "no-cache")
	headers.Set("Connection", "keep-alive")
	headers.Set("X-Accel-Buffering", "no")

	f.Flush()

	for batch := range musicChunks {
		_, err = fmt.Fprintf(w, "data: ")
		_, err = w.Write(batch)
		_, err = fmt.Fprintf(w, "\n\n")
		f.Flush()
	}

	f.Flush()

	return nil
}
