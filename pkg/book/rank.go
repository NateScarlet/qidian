package book

import (
	"bytes"
	"context"
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
	URL url.URL
	// ColumnParser for result table. Defaults to DefaultColumnParser.
	ColumnParser ColumnParser
	Site         string
}

func httpsURL(path string) (ret url.URL) {
	ret.Scheme = "https"
	ret.Host = "www.qidian.com"
	ret.Path = path
	return
}

// Rank types
var (
	// support query by year and month after 2020-01.
	RTMonthlyTicket = RankType{
		URL: httpsURL("/rank/yuepiao/"),
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch strings.TrimSpace(th.Text()) {
			case "":
				book.MonthTicketCount, err = parseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	// support query by year and month after 2020-01.
	RTMonthlyTicketVIP = RankType{
		URL: httpsURL("/rank/fengyun/"),
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch strings.TrimSpace(th.Text()) {
			case "月票榜", "起点月票榜":
				book.MonthTicketCount, err = parseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	// support query by year and month after 2020-01.
	RTMonthlyTicketMM = RankType{
		URL:          httpsURL("/rank/mm/yuepiao/"),
		ColumnParser: RTMonthlyTicketVIP.ColumnParser,
		Site:         "mm",
	}
	RTNewBookSalesMM = RankType{
		URL:  httpsURL("/rank/mm/newsales/"),
		Site: "mm",
	}
	RTDailySales = RankType{
		URL: httpsURL("/rank/hotsales/"),
	}
	RTDailySalesMM = RankType{
		URL:  httpsURL("/rank/mm/hotsales/"),
		Site: "mm",
	}
	RTWeeklyRead = RankType{
		URL: httpsURL("/rank/readindex/"),
	}
	RTWeeklyReadMM = RankType{
		URL:  httpsURL("/rank/mm/readindex/"),
		Site: "mm",
	}
	RTWeeklyRecommendation = RankType{
		URL: httpsURL("/rank/recom/"),
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch column := strings.TrimSpace(th.Text()); column {
			case "推荐":
				book.WeekRecommendCount, err = parseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	RTWeeklyRecommendationMM = RankType{
		URL:          httpsURL("/rank/mm/recom/"),
		ColumnParser: RTWeeklyRecommendation.ColumnParser,
		Site:         "mm",
	}
	RTMonthlyRecommendation = RankType{
		URL: httpsURL("/rank/recom/datatype2/"),
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch strings.TrimSpace(th.Text()) {
			case "推荐":
				book.MonthRecommendCount, err = parseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	RTMonthlyRecommendationMM = RankType{
		URL:          httpsURL("/rank/mm/recom/datatype2/"),
		ColumnParser: RTMonthlyRecommendation.ColumnParser,
		Site:         "mm",
	}
	RTTotalRecommendation = RankType{
		URL: httpsURL("/rank/recom/datatype3/"),
		ColumnParser: ColumnParserFunc(func(book *Book, i int, th, td *goquery.Selection) (err error) {
			switch strings.TrimSpace(th.Text()) {
			case "推荐":
				book.TotalRecommendCount, err = parseSelectionCount(td)
			default:
				err = DefaultColumnParser.ParseColumn(book, i, th, td)
			}
			return
		}),
	}
	RTTotalRecommendationMM = RankType{
		URL:          httpsURL("/rank/mm/recom/datatype3/"),
		ColumnParser: RTTotalRecommendation.ColumnParser,
		Site:         "mm",
	}
	RTTotalBookmark = RankType{
		URL: httpsURL("/rank/collect/"),
	}
	RTTotalBookmarkMM = RankType{
		URL:  httpsURL("/rank/mm/collect/"),
		Site: "mm",
	}
	RTSignedAuthorNewBook = RankType{
		URL: httpsURL("/rank/signnewbook/"),
	}
	RTSignedAuthorNewBookMM = RankType{
		URL:  httpsURL("/rank/mm/signnewbook/"),
		Site: "mm",
	}
	RTPublicAuthorNewBook = RankType{
		URL: httpsURL("/rank/pubnewbook/"),
	}
	RTPublicAuthorNewBookMM = RankType{
		URL:  httpsURL("/rank/mm/pubnewbook/"),
		Site: "mm",
	}
	RTNewSignedAuthorNewBook = RankType{
		URL: httpsURL("/rank/newsign/"),
	}
	RTNewSignedAuthorNewBookMM = RankType{
		URL:  httpsURL("/rank/mm/newsign/"),
		Site: "mm",
	}
	RTNewAuthorNewBook = RankType{
		URL: httpsURL("/rank/newauthor/"),
	}
	RTNewAuthorNewBookMM = RankType{
		URL:  httpsURL("/rank/mm/newauthor/"),
		Site: "mm",
	}
	RTWeeklyFans = RankType{
		URL: httpsURL("/rank/newfans/"),
	}
	RTWeeklyFansMM = RankType{
		URL:  httpsURL("/rank/mm/newfans/"),
		Site: "mm",
	}
	RTLastUpdatedVIP = RankType{
		URL: httpsURL("/rank/vipup/"),
	}
	RTDailyMostUpdateVIPMM = RankType{
		URL:  httpsURL("/rank/mm/vipup/"),
		Site: "mm",
	}
	RTWeeklyMostUpdateVIPMM = RankType{
		URL:  httpsURL("/rank/mm/vipup/datatype2/"),
		Site: "mm",
	}
	RTMonthlyMostUpdateVIPMM = RankType{
		URL:  httpsURL("/rank/mm/vipup/datatype3/"),
		Site: "mm",
	}
	RTTotalWordCountMM = RankType{
		URL:  httpsURL("/rank/mm/wordcount/"),
		Site: "mm",
	}
	RTTotalBookmarkVIP = RankType{
		URL: httpsURL("/rank/vipcollect/"),
	}
	RTWeeklySingleChapterSalesMM = RankType{
		URL:  httpsURL("/rank/mm/subscr/"),
		Site: "mm",
	}
	RTTotalSingleChapterSalesVIPMM = RankType{
		URL:  httpsURL("/rank/mm/vipsub/"),
		Site: "mm",
	}
)

type RankOptions struct {
	category Category
	year     int
	month    time.Month
	page     int
}

type RankOption = func(o *RankOptions)

func RankOptionCategory(v Category) RankOption {
	return func(o *RankOptions) {
		o.category = v
	}
}

// RankOptionYearMonth set rank period, not all rank support this.
func RankOptionYearMonth(year int, month time.Month) RankOption {
	return func(o *RankOptions) {
		o.year = year
		o.month = month
	}
}

// RankOptionMonth set rank period, not all rank support this.
func RankOptionMonth(month time.Month) RankOption {
	return func(o *RankOptions) {
		o.month = month
	}
}

// RankOptionPage set wanted page, start from 1.
func RankOptionPage(page int) RankOption {
	return func(o *RankOptions) {
		o.page = page
	}
}

func RankURL(rt RankType, opts ...RankOption) (ret url.URL) {
	var args = new(RankOptions)
	for _, i := range opts {
		i(args)
	}
	ret = rt.URL
	if !strings.HasSuffix(ret.Path, "/") {
		ret.Path += "/"
	}
	if args.category != "" {
		ret.Path += fmt.Sprintf("chn%s/", string(args.category))
	}
	if args.year != 0 {
		ret.Path += fmt.Sprintf("year%d-month%02d/", args.year, args.month)
	} else if args.month != 0 {
		ret.Path += fmt.Sprintf("month%02d/", args.month)
	}
	if args.page > 1 {
		ret.Path += fmt.Sprintf("page%d/", args.page)
	}
	return
}

type getHTMLResult = client.GetHTMLResult
type RankResult struct {
	getHTMLResult
	rankType RankType
}

func (r RankResult) Books() ([]Book, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(r.Body()))
	if err != nil {
		return nil, err
	}
	table := doc.Find("table.rank-table-list")
	if table.Length() == 0 {
		return nil, fmt.Errorf("qidian: rank: can not found result table: %s", r.Request().URL)
	}
	return parseTable(table, r.RankType().ColumnParser, r.RankType().Site)
}

func Rank(ctx context.Context, rt RankType, opts ...RankOption) (ret RankResult, err error) {
	ret.rankType = rt
	u := RankURL(rt, opts...)
	ret.getHTMLResult, err = client.GetHTML(ctx, u.String(), client.GetHTMLOptionVisitRequest(func(req *http.Request) {
		req.AddCookie(&http.Cookie{
			Name:  "listStyle",
			Value: "2",
		})
	}))
	if err != nil {
		return
	}
	return
}

func (obj RankResult) RankType() RankType {
	return obj.rankType
}
