package wrapper

import (
	"golang.org/x/net/context"
)

type middleware func(HandlerFunc) HandlerFunc
type HandlerFunc func(ctx context.Context, args ...interface{}) (interface{}, error)

func BuildChain(f HandlerFunc, m ...middleware) HandlerFunc {
	if len(m) == 0 {
		return f
	}

	return m[0](BuildChain(f, m[1:cap(m)]...))
}

func (f HandlerFunc) ToWorkerFunc(ctx context.Context, args ...interface{}) func() {
	return func() {
		_, _ = f(ctx, args)

	}
}
