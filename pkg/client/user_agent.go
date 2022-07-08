package client

import "context"

var DefaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0"

type contextKeyUserAgent struct{}

func WithUserAgent(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, contextKeyUserAgent{}, v)
}

func ContextUserAgent(ctx context.Context) string {
	var ret, ok = ctx.Value(contextKeyUserAgent{}).(string)
	if ok {
		return ret
	}
	return DefaultUserAgent
}
