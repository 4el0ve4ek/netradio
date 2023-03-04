package broadcast

import (
	"encoding/json"
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

func (h *getterHandler) ServeJSON(context context.Context) (json.RawMessage, error) {
	return nil, nil
}
