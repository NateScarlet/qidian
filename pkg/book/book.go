package book

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/NateScarlet/qidian/pkg/author"
	"github.com/NateScarlet/qidian/pkg/client"
	"github.com/PuerkitoBio/goquery"
)

// Book model
type Book struct {
	ID string
	// "" for main site,  "mm" for female site
	Site   string
	Title  string
	Author author.Author
	// short description
	Summary string
	// long description
	Introduction string
	Category     Category
	SubCategory  SubCategory
	Tags         []string
	LastUpdated  time.Time
	Finished     time.Time
	WordCount    uint64
	// only avaliable when search by bookmark
	BookmarkCount       uint64
	MonthTicketCount    uint64
	WeekRecommendCount  uint64
	MonthRecommendCount uint64
	TotalRecommendCount uint64
}

// URL of book info page on website.
func (b Book) URL() string {
	return "https://book.qidian.com/info/" + b.ID
}

// Fetch book from info page.
func (b *Book) Fetch(ctx context.Context) (err error) {
	if b.ID == "" {
		return errors.New("empty book id")
	}

	req, err := http.NewRequestWithContext(ctx, "GET", b.URL(), nil)
	if err != nil {
		return err
	}
	resp, err := client.For(ctx).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	introElem := doc.Find(".book-info").Clone()
	stateElem := doc.Find(".book-state")

	// Author
	writerElem := introElem.Find("a.writer")
	writerElem.Parent().Remove()
	b.Author.Name = writerElem.Text()
	if href := writerElem.AttrOr("href", ""); strings.HasPrefix(href, "//me.qidian.com/authorIndex.aspx?id=") {
		b.Author.ID = href[36:]
	}

	// Title
	b.Title = strings.TrimSpace(introElem.Find("h1").Text())

	// Categories
	introElem.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}
		u, err := url.Parse(href)
		if err != nil {
			return
		}
		if s := u.Query().Get("chanId"); s != "" {
			b.Category = Category(s)
		}
		if s := u.Query().Get("subCateId"); s != "" {
			b.SubCategory = SubCategory(s)
		}
	})

	// Tags
	tagElemList := introElem.Find(".tag > span").
		AddSelection(stateElem.Find(".tags"))
	b.Tags = make([]string, 0, tagElemList.Length())
	tagElemList.Each(func(i int, s *goquery.Selection) {
		b.Tags = append(b.Tags, s.Text())
	})

	// Introduction
	b.Summary = introElem.Find(".intro").Text()
	b.Introduction = nodesText(doc.Find(".book-info-detail .book-intro").Nodes)

	// Count
	introElem.Find("style").EachWithBreak(func(i int, s *goquery.Selection) bool {
		var parent = s.Parent()
		var c string
		c, err = deobfuscate(parent)
		if err != nil {
			return false
		}
		c += parent.Next().Text()
		if strings.HasSuffix(c, "字") {
			b.WordCount, err = ParseCount(c[:len(c)-len("字")])
			if err != nil {
				return false
			}
		} else if strings.HasSuffix(c, "总推荐") {
			b.TotalRecommendCount, err = ParseCount(c[:len(c)-len("总推荐")])
			if err != nil {
				return false
			}
		} else if strings.HasSuffix(c, "周推荐") {
			b.WeekRecommendCount, err = ParseCount(c[:len(c)-len("周推荐")])
			if err != nil {
				return false
			}
		}
		return true
	})
	if err != nil {
		return err
	}

	// LastUpdated
	b.LastUpdated, err = ParseTime(stateElem.Find(".update .time").Text())
	if err != nil {
		return err
	}

	// MonthTicket
	b.MonthTicketCount, err = strconv.ParseUint(doc.Find("#monthCount").Text(), 10, 64)
	if err != nil {
		return err
	}

	return nil
}
