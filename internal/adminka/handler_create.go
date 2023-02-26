package adminka

import (
	"encoding/json"
	"net/http"
	"netradio/internal/databases/news"
	"netradio/libs/context"
	"netradio/models"
	"netradio/pkg/handles"
	"time"
)

func newCreateHandler(newsService news.Service) *createHandler {
	return &createHandler{
		newsService: newsService,
	}
}

type createHandler struct {
	newsService news.Service
}

func (h *createHandler) ServeHTTP(context context.Context, request *http.Request) (handles.Response, error) {
	var rawNews struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&rawNews)
	if err != nil {
		return handles.Response{}, err
	}

	newNews := models.News{
		Title:           rawNews.Title,
		Content:         rawNews.Content,
		PublicationTime: time.Now().Unix(),
	}

	h.newsService.Add(newNews)

	return handles.Response{}, nil
}
