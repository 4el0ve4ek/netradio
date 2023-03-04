package news

import (
	"encoding/json"
	"netradio/internal/databases/news"
	"netradio/libs/context"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func newGetterHandler(newsService news.Service) *getterHandler {
	return &getterHandler{
		newsService: newsService,
	}
}

type getterHandler struct {
	newsService news.Service
}

func (h *getterHandler) ServeJSON(context context.Context) (json.RawMessage, error) {
	newsID := chi.URLParam(context.GetRequest(), "newsID")
	id, err := strconv.Atoi(newsID)
	if err != nil {
		return nil, err
	}

	newsOne, err := h.newsService.GetByID(id)
	if err != nil {
		return nil, err
	}

	res, err := json.Marshal(newsOne)
	if err != nil {
		return nil, err
	}
	return res, nil
}
