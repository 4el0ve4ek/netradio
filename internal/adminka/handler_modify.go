package adminka

import (
	"encoding/json"
	"errors"
	"net/http"
	"netradio/internal/databases/news"
	"netradio/libs/context"
	"netradio/models"
	"netradio/pkg/handles"
	"time"
)

func newModifyHandler(newsService news.Service) *modifyHandler {
	return &modifyHandler{
		newsService: newsService,
	}
}

type modifyHandler struct {
	newsService news.Service
}

func (h *modifyHandler) ServeHTTP(context context.Context, request *http.Request) (handles.Response, error) {
	var rawNews struct {
		ID      *int    `json:"id,omitempty"`
		Title   *string `json:"title,omitempty"`
		Content *string `json:"content,omitempty"`

		//PublicationTime *int64  `json:"publication_date,omitempty"` // <- not allowed to modify
	}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&rawNews)

	if err != nil {
		return handles.Response{}, err
	}
	if rawNews.ID == nil {
		return handles.Response{}, errors.New("no such have been passed")
	}

	modifiedNews, err := h.newsService.GetByID(*rawNews.ID)
	if err != nil {
		modifiedNews = models.News{
			ID: *rawNews.ID,
		}
	}

	if rawNews.Title != nil {
		modifiedNews.Title = *rawNews.Title
	}

	if rawNews.Content != nil {
		modifiedNews.Content = *rawNews.Content
	}

	modifiedNews.PublicationTime = time.Now().Unix()
	h.newsService.Add(modifiedNews)
	return handles.Response{}, nil
}
