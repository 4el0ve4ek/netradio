package handles

import (
	"encoding/json"
	"net/http"
	"netradio/libs/context"
)

type Handler interface {
	ServeHTTP(context.Context, *http.Request) (Response, error)
}

type HandlerJSON interface {
	ServeJSON(context context.Context, r *http.Request) (json.RawMessage, error)
}
