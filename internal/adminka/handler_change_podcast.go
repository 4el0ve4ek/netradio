package adminka

import (
	"io"
	"netradio/internal/databases/music"
	"netradio/libs/context"
	"netradio/pkg/handles"
	"os"
)

func newChangePodcastHandler(musicServise music.Service) *changePodcastHandler {
	return &changePodcastHandler{
		musicServise: musicServise,
	}
}

type changePodcastHandler struct {
	musicServise music.Service
}

func (h *changePodcastHandler) ServeHTTP(context context.Context) (handles.Response, error) {
	request := context.GetRequest()
	err := request.ParseMultipartForm(maxPodcastSize)
	if err != nil {
		return handles.Response{}, err
	}

	file, _, err := request.FormFile("music_file")
	if err != nil {
		return handles.Response{}, err
	}
	content, err := io.ReadAll(file)
	if err != nil {
		return handles.Response{}, err
	}
	err = os.WriteFile("new_music", content, os.ModePerm)

	request.PostFormValue("id")
	request.PostFormValue("title")
	request.PostFormValue("photo_link") // ?

	if err != nil {
		return handles.Response{}, err
	}

	return handles.Response{}, nil
}
