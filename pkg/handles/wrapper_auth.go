package handles

import (
	"net/http"
	"netradio/libs/context"
)

func NewAuthWrapper(handler HandlerWritable, core Core) *authWrapper {
	return &authWrapper{
		core:     core,
		original: handler,
	}
}

type authWrapper struct {
	core     Core
	original HandlerWritable
}

func (w *authWrapper) ServeHTTP(responseWriter http.ResponseWriter, r *http.Request) {
	ctx := context.New(r, w.core.auth, w.core.userService, w.core.log)
	err := w.original.ServeHTTP(ctx, responseWriter)

	if err != nil {
		ctx.GetLogger().Warn(err)
	}
}
