// Package client configure http client that used for qidian requests.
package client

import (
	"context"
	"net/http"
)

type contextKey struct{}

// For get client from context.
func For(ctx context.Context) *http.Client {
	v := ctx.Value(contextKey{}).(*http.Client)
	if v == nil {
		return http.DefaultClient
	}

	return v
}

// With set client to context.
func With(ctx context.Context, v *http.Client) context.Context {
	return context.WithValue(ctx, contextKey{}, v)
}
