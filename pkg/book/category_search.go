package book

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

// Sort for search
type Sort string

// Sort for search
const (
	STotalRecommend Sort = "2"
	SCharCount           = "3"
	SLastUpdated         = "5"
	SRecentFinished      = "6"
	SWeekRecommend       = "9"
	SMonthRecommend      = "10"
	STotalBookmark       = "11"
)

// CategorySearch use https://www.qidian.com/all page
type CategorySearch struct {
	Site        string
	Sort        Sort
	Page        int
	Category    Category
	SubCategory SubCategory
}

// NewCategorySearch create a new search for function chaining.
func NewCategorySearch() *CategorySearch {
	return &CategorySearch{}
}

// SetPage then returns self.
func (s *CategorySearch) SetPage(v int) *CategorySearch {
	s.Page = v
	return s
}

// SetSort then returns self.
func (s *CategorySearch) SetSort(v Sort) *CategorySearch {
	s.Sort = v
	return s
}

// SetCategory then returns self.
func (s *CategorySearch) SetCategory(v Category) *CategorySearch {
	s.Category = v
	s.Site = v.Site()
	return s
}

// SetSubCategory and category then returns self.
func (s *CategorySearch) SetSubCategory(v SubCategory) *CategorySearch {
	s.SubCategory = v
	s.SetCategory(v.Parent())
	return s
}

// Execute search
func (s CategorySearch) Execute(ctx context.Context) (ret []Book, err error) {
	var base = "https://www.qidian.com"
	if s.Site != "" {
		base += "/" + s.Site
	}
	u, err := url.Parse(base + "/all")
	if err != nil {
		return
	}
	q := u.Query()
	q.Set("style", "2")
	if s.Page > 1 {
		q.Set("page", strconv.Itoa(s.Page))
	}
	if s.Sort != "" {
		q.Set("orderId", string(s.Sort))
	}
	if s.Category != "" {
		q.Set("chanId", string(s.Category))
	}
	if s.SubCategory != "" {
		q.Set("subCateId", string(s.SubCategory))
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	table := doc.
		Find("table.rank-table-list")
	if table.Length() == 0 {
		return nil, errors.New("can not found result table")
	}
	return parseTable(table, nil, s.Site)
}
