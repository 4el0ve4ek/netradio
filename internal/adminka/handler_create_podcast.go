package adminka

import (
	"io"
	"netradio/internal/databases/music"
	"netradio/libs/context"
	"netradio/pkg/handles"
	"os"
)

const maxPodcastSize = 1 << 30 // 1 gigabyte

func newCreatePodcastHandler(musicService music.Service) *createHandler {
	return &createHandler{
		musicService: musicService,
	}
}

type createHandler struct {
	musicService music.Service
}

func (h *createHandler) ServeHTTP(context context.Context) (handles.Response, error) {
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

	//request.PostFormValue("id") // no id when creation yet
	request.PostFormValue("title")
	request.PostFormValue("photo_link") // ?

	if err != nil {
		return handles.Response{}, err
	}

	return handles.Response{}, nil
}
