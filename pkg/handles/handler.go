package handles

import (
	"encoding/json"
	"net/http"
	"netradio/libs/context"
)

type Handler interface {
	ServeHTTP(context.Context) (Response, error)
}

type HandlerWritable interface {
	ServeHTTP(context.Context, http.ResponseWriter) error
}

type HandlerJSON interface {
	ServeJSON(context context.Context) (json.RawMessage, error)
}
