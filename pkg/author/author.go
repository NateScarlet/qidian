package author

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/NateScarlet/qidian/pkg/client"
	"github.com/NateScarlet/qidian/pkg/util"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// Author of book, has different id from user.
type Author struct {
	ID        string
	UserID    string
	Name      string
	AvatarURL string
	Biography string
}

// URL of author info page.
func (a Author) URL() string {
	return "https://my.qidian.com/author/" + a.ID
}

// Fetch data from author info page.
func (a *Author) Fetch(ctx context.Context) (err error) {

	if err != nil {
		return errors.New("qidian: empty author id")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", a.URL(), nil)
	if err != nil {
		return
	}
	resp, err := client.For(ctx).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}

	// name
	if nodes := doc.Find(".header-msg > h3").Nodes; len(nodes) == 1 {
		n := nodes[0].FirstChild
		if n != nil && n.Type == html.TextNode {
			a.Name = n.Data
		}
	}

	// user id
	userHref := doc.Find("a.header-msg-tosingle").AttrOr("href", "")
	if strings.HasPrefix(userHref, "/user/") {
		a.UserID = userHref[6:]
	}

	//
	a.AvatarURL = util.AbsoluteURL(doc.Find("img.header-avatar-img").AttrOr("src", ""))

	//
	a.Biography = strings.Trim(doc.Find(".header-msg-desc").Text(), " \n")

	return
}
