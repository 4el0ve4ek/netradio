package handles

import (
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

func (w *wrapperJSON) ServeHTTP(ctx context.Context) (Response, error) {
	res, err := w.original.ServeJSON(ctx)
	if err != nil {
		return Response{}, err
	}

	return Response{
		Content:    res,
		StatusCode: 200,
	}, nil
}
