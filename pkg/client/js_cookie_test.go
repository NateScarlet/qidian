package client

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJSCookieSrc(t *testing.T) {
	var ctx = context.Background()
	var doc *goquery.Document
	func() {
		r, err := os.Open("js_cookie_sample.html")
		require.NoError(t, err)
		defer r.Close()
		doc, err = goquery.NewDocumentFromReader(r)
		require.NoError(t, err)
	}()

	data, err := newJSCookieTemplateData("https://book.qidian.com/info/1004608738/", *doc)
	require.NoError(t, err)
	url, err := jsCookieSrc(ctx, *data)
	require.NoError(t, err)
	assert.Equal(t, "https://book.qidian.com/b3c79ec/f890b6f5917/53f27290.js", url)
}

func TestJSCookieValue(t *testing.T) {
	var ctx = context.Background()
	var doc *goquery.Document
	func() {
		r, err := os.Open("js_cookie_sample.html")
		require.NoError(t, err)
		defer r.Close()
		doc, err = goquery.NewDocumentFromReader(r)
		require.NoError(t, err)
	}()

	data, err := newJSCookieTemplateData("https://book.qidian.com/info/1004608738/", *doc)
	require.NoError(t, err)
	script2Data, err := ioutil.ReadFile("js_cookie_sample.js")
	require.NoError(t, err)
	data.Script2 = string(script2Data)

	cookie, err := jsCookieValue(ctx, *data)
	require.NoError(t, err)
	assert.Regexp(t, "^Cc2838679FT=.{173}; path=/; expires=.{29}$", cookie)
}
