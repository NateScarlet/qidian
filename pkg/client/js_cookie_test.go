package client

import (
	"context"
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

	url, err := jsCookieSrc(ctx, "https://book.qidian.com/info/1004608738/", *doc)
	require.NoError(t, err)
	assert.Equal(t, "https://book.qidian.com/b3c79ec/f890b6f5917/53f27290.js", url)
}
