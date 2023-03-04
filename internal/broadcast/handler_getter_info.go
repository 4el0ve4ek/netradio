package broadcast

import (
	"encoding/json"
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

func (h *getterInfoHandler) ServeJSON(context context.Context) (json.RawMessage, error) {
	return nil, nil
}
