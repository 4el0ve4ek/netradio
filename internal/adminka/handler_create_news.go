package adminka

import (
	"encoding/json"
	"netradio/internal/databases/news"
	"netradio/libs/context"
	"netradio/models"
	"netradio/pkg/handles"
	"time"
)

func newCreateNewsHandler(newsService news.Service) *createNewsHandler {
	return &createNewsHandler{
		newsService: newsService,
	}
}

type createNewsHandler struct {
	newsService news.Service
}

func (h *createNewsHandler) ServeHTTP(context context.Context) (handles.Response, error) {
	request := context.GetRequest()
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
