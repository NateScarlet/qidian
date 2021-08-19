package book

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/NateScarlet/qidian/pkg/client"
	"github.com/PuerkitoBio/goquery"
)

// RankType that corresponding a rank page.
type RankType struct {
	// URL for rank page.
	URL string
	// ColumnParser for result table. Defaults to DefaultColumnParser.
	ColumnParser ColumnParser
	Site         string
}

// Rank types
var (
	// support query by year and month after 2020-01.
	RTMonthlyTicket = RankType{
		URL: "https://www.qidian.com/rank/yuepiao/",
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch strings.TrimSpace(th.Text()) {
			case "":
				book.MonthTicketCount, err = ParseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	// support query by year and month after 2020-01.
	RTMonthlyTicketVIP = RankType{
		URL: "https://www.qidian.com/rank/fengyun/",
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch strings.TrimSpace(th.Text()) {
			case "月票榜", "起点月票榜":
				book.MonthTicketCount, err = ParseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	// support query by year and month after 2020-01.
	RTMonthlyTicketMM = RankType{
		URL:          "https://www.qidian.com/rank/mm/yuepiao/",
		ColumnParser: RTMonthlyTicketVIP.ColumnParser,
		Site:         "mm",
	}
	RTNewBookSalesMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/newsales/",
		Site: "mm",
	}
	RTDailySales = RankType{
		URL: "https://www.qidian.com/rank/hotsales/",
	}
	RTDailySalesMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/hotsales/",
		Site: "mm",
	}
	RTWeeklyRead = RankType{
		URL: "https://www.qidian.com/rank/readindex/",
	}
	RTWeeklyReadMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/readindex/",
		Site: "mm",
	}
	RTWeeklyRecommendation = RankType{
		URL: "https://www.qidian.com/rank/recom/",
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch column := strings.TrimSpace(th.Text()); column {
			case "推荐":
				book.WeekRecommendCount, err = ParseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	RTWeeklyRecommendationMM = RankType{
		URL:          "https://www.qidian.com/rank/mm/recom/",
		ColumnParser: RTWeeklyRecommendation.ColumnParser,
		Site:         "mm",
	}
	RTMonthlyRecommendation = RankType{
		URL: "https://www.qidian.com/rank/recom/datatype2/",
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch strings.TrimSpace(th.Text()) {
			case "推荐":
				book.MonthRecommendCount, err = ParseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	RTMonthlyRecommendationMM = RankType{
		URL:          "https://www.qidian.com/rank/mm/recom/datatype2/",
		ColumnParser: RTMonthlyRecommendation.ColumnParser,
		Site:         "mm",
	}
	RTTotalRecommendation = RankType{
		URL: "https://www.qidian.com/rank/recom/datatype3/",
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch strings.TrimSpace(th.Text()) {
			case "推荐":
				book.TotalRecommendCount, err = ParseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	RTTotalRecommendationMM = RankType{
		URL:          "https://www.qidian.com/rank/mm/recom/datatype3/",
		ColumnParser: RTTotalRecommendation.ColumnParser,
		Site:         "mm",
	}
	RTTotalBookmark = RankType{
		URL: "https://www.qidian.com/rank/collect/",
	}
	RTTotalBookmarkMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/collect/",
		Site: "mm",
	}
	RTSignedAuthorNewBook = RankType{
		URL: "https://www.qidian.com/rank/signnewbook/",
	}
	RTSignedAuthorNewBookMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/signnewbook/",
		Site: "mm",
	}
	RTPublicAuthorNewBook = RankType{
		URL: "https://www.qidian.com/rank/pubnewbook/",
	}
	RTPublicAuthorNewBookMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/pubnewbook/",
		Site: "mm",
	}
	RTNewSignedAuthorNewBook = RankType{
		URL: "https://www.qidian.com/rank/newsign/",
	}
	RTNewSignedAuthorNewBookMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/newsign/",
		Site: "mm",
	}
	RTNewAuthorNewBook = RankType{
		URL: "https://www.qidian.com/rank/newauthor/",
	}
	RTNewAuthorNewBookMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/newauthor/",
		Site: "mm",
	}
	RTWeeklyFans = RankType{
		URL: "https://www.qidian.com/rank/newfans/",
	}
	RTWeeklyFansMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/newfans/",
		Site: "mm",
	}
	RTLastUpdatedVIP = RankType{
		URL: "https://www.qidian.com/rank/vipup/",
	}
	RTDailyMostUpdateVIPMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/vipup/",
		Site: "mm",
	}
	RTWeeklyMostUpdateVIPMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/vipup/datatype2/",
		Site: "mm",
	}
	RTMonthlyMostUpdateVIPMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/vipup/datatype3/",
		Site: "mm",
	}
	RTTotalWordCountMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/wordcount/",
		Site: "mm",
	}
	RTTotalBookmarkVIP = RankType{
		URL: "https://www.qidian.com/rank/vipcollect/",
	}
	RTWeeklyRewardVIP = RankType{
		URL: "https://www.qidian.com/rank/vipreward/",
	}
	RTWeeklySingleChapterSalesMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/subscr/",
		Site: "mm",
	}
	RTTotalSingleChapterSalesVIPMM = RankType{
		URL:  "https://www.qidian.com/rank/mm/vipsub/",
		Site: "mm",
	}
)

// Rank for books.
// not all rank type support query by year and month.
type Rank struct {
	Type     RankType
	Category Category
	Year     int
	Month    time.Month
}

// URL of rank page.
func (r Rank) URL() string {
	var u, err = url.Parse(r.Type.URL)
	if err != nil {
		return ""
	}
	if !strings.HasSuffix(u.Path, "/") {
		u.Path += "/"
	}
	if r.Category != "" {
		u.Path += fmt.Sprintf("chn%s/", string(r.Category))
	}
	if r.Year != 0 {
		u.Path += fmt.Sprintf("year%d-month%02d/", r.Year, r.Month)
	} else if r.Month != 0 {
		u.Path += fmt.Sprintf("month%02d/", r.Month)
	}
	return u.String()
}

// Fetch rank, return 50 book in order.
func (r Rank) Fetch(ctx context.Context) ([]Book, error) {
	var u = r.URL()
	if u == "" {
		return nil, errors.New("qidian: invalid rank url")
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{
		Name:  "listStyle",
		Value: "2",
	})
	resp, err := client.For(ctx).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	table := doc.Find("table.rank-table-list")
	if table.Length() == 0 {
		return nil, fmt.Errorf("qidian: rank: can not found result table: %s", u)
	}
	return parseTable(table, r.Type.ColumnParser, r.Type.Site)
}
