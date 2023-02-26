package news

import (
	"encoding/json"
	"net/http"
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

func (h *getterHandler) ServeJSON(context context.Context, request *http.Request) (json.RawMessage, error) {
	newsID := chi.URLParam(request, "newsID")
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
