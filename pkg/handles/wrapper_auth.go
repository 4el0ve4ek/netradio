package handles

import (
	"net/http"
	"netradio/libs/context"
)

func NewAuthWrapper(handler Handler, core Core) *authWrapper {
	return &authWrapper{
		core:     core,
		original: handler,
	}
}

type authWrapper struct {
	core     Core
	original Handler
}

func (w *authWrapper) ServeHTTP(responseWriter http.ResponseWriter, r *http.Request) {
	ctx := context.New(r, w.core.auth, w.core.log)
	response, err := w.original.ServeHTTP(ctx, r)

	for k, v := range response.Headers {
		if len(v) == 0 {
			continue
		}
		responseWriter.Header().Add(k, v[0])
	}

	responseWriter.Header().Add("Access-Control-Allow-Origin", "*")

	var statusCode int
	if err != nil {
		ctx.GetLogger().Warn(err)
		statusCode = http.StatusInternalServerError
	} else {
		statusCode = response.GetStatusCodeOrDefault(http.StatusOK)
	}
	responseWriter.WriteHeader(statusCode)

	if response.Content == nil {
		return
	}
	_, err = responseWriter.Write(response.GetContent())
	if err != nil {
		ctx.GetLogger().Warn(err)
	}
}
