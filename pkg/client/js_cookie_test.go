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

	_, url, err := jsCookieSrc(ctx, "https://book.qidian.com/info/1004608738/", *doc)
	require.NoError(t, err)
	assert.Equal(t, "https://book.qidian.com/b3c79ec/f890b6f5917/53f27290.js", url)
}

func TestJSCookieTODO(t *testing.T) {
	var ctx = context.Background()
	var doc *goquery.Document
	func() {
		r, err := os.Open("js_cookie_sample.html")
		require.NoError(t, err)
		defer r.Close()
		doc, err = goquery.NewDocumentFromReader(r)
		require.NoError(t, err)
	}()

	script1, _, err := jsCookieSrc(ctx, "https://book.qidian.com/info/1004608738/", *doc)
	require.NoError(t, err)
	script2, err := ioutil.ReadFile("53f27290_v1.local.js")
	require.NoError(t, err)

	cookie, err := jsCookieTODO(ctx, "https://book.qidian.com/info/1004608738/", script1, string(script2))
	require.NoError(t, err)
	assert.Equal(t, "Cc2838679FT=637r_mbhoxLVqqqDrBeuv6GSwOE3JJrUxnpmSQijotp0e6RrRDwrrX72PompklXYQ0gWXAVUKLZU2DjBRFH3cl26JzdrfbRouCsrbv.M7rfEGPZJ769dPXxDiWVGU9_dKIpwxVOxtKrin1TpyX_MAwVMnzfo5lZF5XjmQuqHg_Qma; path=/; expires=Thu, 14 Jul 2022 18:31:33 GMT", cookie)
}

// func TestJSCookieTODO2(t *testing.T) {
// 	var ctx = context.Background()
// 	var doc *goquery.Document
// 	func() {
// 		r, err := os.Open("js_cookie_sample.html")
// 		require.NoError(t, err)
// 		defer r.Close()
// 		doc, err = goquery.NewDocumentFromReader(r)
// 		require.NoError(t, err)
// 	}()

// 	script1, _, err := jsCookieSrc(ctx, "https://book.qidian.com/info/1004608738/", *doc)
// 	require.NoError(t, err)
// 	script2, err := ioutil.ReadFile("53f27290_v1.local.js")
// 	require.NoError(t, err)
// 	script3, err := ioutil.ReadFile("script3_sample_v1.local.js")
// 	require.NoError(t, err)

// 	url, err := jsCookieTODO2(ctx, "https://book.qidian.com/info/1004608738/", script1, string(script2), string(script3))
// 	require.NoError(t, err)
// 	assert.Equal(t, "https://book.qidian.com/b3c79ec/f890b6f5917/53f27290.js", url)
// }
