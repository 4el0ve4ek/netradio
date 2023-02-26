package broadcast

import (
	"encoding/json"
	"net/http"
	"netradio/internal/databases/channel"
	"netradio/libs/context"
)

func newInfoGetterHandler(channelService channel.Service) *getterInfoHandler {
	return &getterInfoHandler{
		channelService: channelService,
	}
}

type getterInfoHandler struct {
	channelService channel.Service
}

func (h *getterInfoHandler) ServeJSON(context context.Context, r *http.Request) (json.RawMessage, error) {
	return nil, nil
}
