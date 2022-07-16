package book

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/NateScarlet/qidian/pkg/author"
	"github.com/NateScarlet/qidian/pkg/client"
	"github.com/NateScarlet/qidian/pkg/util"
	"github.com/PuerkitoBio/goquery"
)

// Book model
type Book struct {
	ID string
	// "" for main site,  "mm" for female site
	Site     string
	Title    string
	Author   author.Author
	CoverURL string
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
	// only available when search by bookmark
	BookmarkCount       uint64
	MonthTicketCount    uint64
	WeekRecommendCount  uint64
	MonthRecommendCount uint64
	TotalRecommendCount uint64
}

// URL of book info page on website.
func (b Book) URL() string {
	return "https://book.qidian.com/info/" + b.ID + "/"
}

var categoryURLPattern = regexp.MustCompile(`//www\.qidian\.com/all/chanId(\d+)-subCateId(\d+)/`)
var authorURLPattern = regexp.MustCompile(`//my\.qidian\.com/author/(\d+)/`)

// Fetch book from info page.
func (b *Book) Fetch(ctx context.Context) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("qidian: (*Book{'%s'}).Fetch: %w", b.ID, err)
		}
	}()

	if b.ID == "" {
		return errors.New("empty book id")
	}

	var url = b.URL()
	getHTML, err := client.GetHTML(ctx, url)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(getHTML.Body()))
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			html, _ := doc.Html()
			err = fmt.Errorf("document=%s:\n%w", html, err)
		}
	}()
	infoElem := doc.Find(".book-info").Clone()
	stateElem := doc.Find(".book-state")

	// Author
	writerElem := infoElem.Find("a.writer")
	writerElem.Parent().Remove()
	b.Author.Name = writerElem.Text()
	if href := writerElem.AttrOr("href", ""); href != "" {
		var match = authorURLPattern.FindStringSubmatch(href)
		if len(match) == 2 {
			b.Author.ID = match[1]
		}
	}

	// Title
	b.Title = strings.TrimSpace(infoElem.Find("h1").Text())

	// Categories
	infoElem.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}
		for _, match := range categoryURLPattern.FindAllStringSubmatch(href, -1) {
			b.Category = Category(match[1])
			b.SubCategory = SubCategory(match[2])
		}
	})

	// Cover
	b.CoverURL = util.AbsoluteURL(doc.Find("#bookImg > img").AttrOr("src", ""))

	// Tags
	tagElemList := infoElem.Find(".tag > span").
		AddSelection(stateElem.Find(".tags"))
	b.Tags = make([]string, 0, tagElemList.Length())
	tagElemList.Each(func(i int, s *goquery.Selection) {
		b.Tags = append(b.Tags, s.Text())
	})

	// Introduction
	b.Summary = infoElem.Find(".intro").Text()
	b.Introduction = nodesText(doc.Find(".book-info-detail .book-intro").Nodes)

	// Count
	infoElem.Find("style").EachWithBreak(func(i int, s *goquery.Selection) bool {
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
	b.LastUpdated, err = ParseTime(stateElem.Find(".update .time").First().Text())
	if err != nil {
		return err
	}

	// MonthTicket
	monthlyTickerEl := doc.Find("#monthCount")
	if monthlyTickerEl.Length() > 0 {
		b.MonthTicketCount, err = strconv.ParseUint(monthlyTickerEl.Text(), 10, 64)
		if err != nil {
			return err
		}
	}

	return nil
}
