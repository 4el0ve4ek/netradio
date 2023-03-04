package news

import (
	"encoding/json"
	"netradio/internal/databases/news"
	"netradio/libs/context"
	"sort"
)

const (
	defaultRecentLen = 10
)

func newRecentGetterHandler(newsService news.Service) *recentHandler {
	return &recentHandler{
		newsService: newsService,
	}
}

type recentHandler struct {
	newsService news.Service
}

func (h *recentHandler) ServeJSON(context context.Context) (json.RawMessage, error) {
	newsAll := h.newsService.GetAll()

	sort.Slice(newsAll, func(i, j int) bool {
		return newsAll[i].PublicationTime > newsAll[j].PublicationTime
	})

	res, err := json.Marshal(newsAll)
	if err != nil {
		return nil, err
	}
	return res, nil
}
