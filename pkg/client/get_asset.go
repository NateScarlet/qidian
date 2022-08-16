package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetAsset(ctx context.Context, url string) (body []byte, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("GetAsset('%s'): %w", url, err)
		}
	}()
	var cache = ContextAssetCache(ctx)
	body, ok, err := cache.Get(ctx, url)
	if err != nil {
		return
	}
	if ok {
		return
	}
	req, err := newRequest(ctx, "GET", url, nil)
	if err != nil {
		return
	}
	resp, err := For(ctx).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status %d", resp.StatusCode)
		return
	}
	if strings.HasSuffix(url, ".js") && !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/javascript") {
		err = fmt.Errorf("unexpected content type '%s'", resp.Header.Get("Content-Type"))
		return
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = cache.Set(ctx, url, body)
	if err != nil {
		return
	}
	return
}
