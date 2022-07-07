// Package client configure http client that used for qidian requests.
package client

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"os"
)

var isDebug = os.Getenv("DEBUG") == "qidian.client"

type contextKey struct{}

var Default = new(http.Client)

func init() {
	Default.Jar, _ = cookiejar.New(nil)
}

// For get client from context.
func For(ctx context.Context) *http.Client {
	v, _ := ctx.Value(contextKey{}).(*http.Client)
	if v == nil {
		return Default
	}
	return v
}

// With set client to context.
func With(ctx context.Context, v *http.Client) context.Context {
	return context.WithValue(ctx, contextKey{}, v)
}
