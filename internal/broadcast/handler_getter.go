package broadcast

import (
	"encoding/json"
	"net/http"
	"netradio/internal/databases/channel"
	"netradio/libs/context"
)

func newGetterHandler(channelService channel.Service) *getterHandler {
	return &getterHandler{
		channelService: channelService,
	}
}

type getterHandler struct {
	channelService channel.Service
}

func (h *getterHandler) ServeJSON(context context.Context, r *http.Request) (json.RawMessage, error) {
	return nil, nil
}
