package handles

import (
	"net/http"
	"netradio/libs/context"
)

func NewJSONWrapper(handler HandlerJSON) *wrapperJSON {
	return &wrapperJSON{
		original: handler,
	}
}

type wrapperJSON struct {
	original HandlerJSON
}

func (w *wrapperJSON) ServeHTTP(ctx context.Context, r *http.Request) (Response, error) {
	res, err := w.original.ServeJSON(ctx, r)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Content:    res,
		StatusCode: 200,
	}, nil
}
