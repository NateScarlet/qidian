package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func handleJSProtect(ctx context.Context, res *GetHTMLResult) (err error) {
	if !bytes.Contains(res.Body(), []byte("_rspj")) {
		return
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res.Body()))
	if err != nil {
		return
	}
	data, err := newJSCookieTemplateData(ctx, res.Request().URL.String(), *doc)
	if err != nil {
		return
	}
	if data == nil {
		return nil
	}
	var c = For(ctx)
	if c.Jar == nil {
		err = fmt.Errorf("nil cookie jar")
		return
	}
	cookie, err := jsCookie(ctx, *data)
	if err != nil {
		return
	}
	c.Jar.SetCookies(res.Request().URL, []*http.Cookie{cookie})
	resp, err := c.Do(res.Request())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	res.response = resp
	res.body, err = io.ReadAll(resp.Body)
	return
}

type GetHTMLOptions struct {
	visitRequest  func(req *http.Request)
	visitResponse func(resp *http.Response)
}

func newGetHTMLOptions(options ...GetHTMLOption) *GetHTMLOptions {
	var opts = new(GetHTMLOptions)
	for _, i := range options {
		i(opts)
	}
	return opts
}

type GetHTMLOption func(opts *GetHTMLOptions)

type GetHTMLResult struct {
	response *http.Response
	body     []byte
}

func (obj GetHTMLResult) Request() *http.Request {
	if obj.response == nil {
		return nil
	}
	return obj.response.Request
}

func (obj GetHTMLResult) Response() *http.Response {
	return obj.response
}

func (obj GetHTMLResult) Body() []byte {
	return obj.body
}

var getHTMLMu sync.Mutex

func GetHTML(ctx context.Context, url string, options ...GetHTMLOption) (res GetHTMLResult, err error) {
	getHTMLMu.Lock()
	defer getHTMLMu.Unlock()
	var opts = newGetHTMLOptions(options...)
	req, err := newRequest(ctx, "GET", url, nil)
	if err != nil {
		return
	}
	if opts.visitRequest != nil {
		opts.visitRequest(req)
	}
	resp, err := For(ctx).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	res.response = resp
	res.body, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = handleJSProtect(ctx, &res)
	if err != nil {
		return
	}
	return
}

func GetHTMLOptionVisitRequest(visitor func(req *http.Request)) GetHTMLOption {
	return func(opts *GetHTMLOptions) {
		opts.visitRequest = visitor
	}
}
