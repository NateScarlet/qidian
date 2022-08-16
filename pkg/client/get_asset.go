package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func GetAsset(ctx context.Context, url string) (body []byte, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("GetAsset('%s'): %w", url, err)
		}
	}()

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
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
