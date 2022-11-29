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

func mainSiteURL(path string) (ret url.URL) {
	ret.Scheme = "https"
	ret.Host = "www.qidian.com"
	ret.Path = path
	return
}

func mmSiteURL(path string) (ret url.URL) {
	ret.Scheme = "https"
	ret.Host = "www.qdmm.com"
	ret.Path = path
	return
}

// Rank types
var (
	// support query by year and month after 2020-01.
	RTMonthlyTicket = RankType{
		URL: mainSiteURL("/rank/yuepiao/"),
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
		URL: mainSiteURL("/rank/fengyun/"),
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
		URL:          mmSiteURL("/rank/yuepiao/"),
		ColumnParser: RTMonthlyTicketVIP.ColumnParser,
		Site:         "mm",
	}
	RTNewBookSalesMM = RankType{
		URL:  mmSiteURL("/rank/newsales/"),
		Site: "mm",
	}
	RTDailySales = RankType{
		URL: mainSiteURL("/rank/hotsales/"),
	}
	RTDailySalesMM = RankType{
		URL:  mmSiteURL("/rank/hotsales/"),
		Site: "mm",
	}
	RTWeeklyRead = RankType{
		URL: mainSiteURL("/rank/readindex/"),
	}
	RTWeeklyReadMM = RankType{
		URL:  mmSiteURL("/rank/readindex/"),
		Site: "mm",
	}
	RTWeeklyRecommendation = RankType{
		URL: mainSiteURL("/rank/recom/"),
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
		URL:          mmSiteURL("/rank/recom/"),
		ColumnParser: RTWeeklyRecommendation.ColumnParser,
		Site:         "mm",
	}
	RTMonthlyRecommendation = RankType{
		URL: mainSiteURL("/rank/recom/datatype2/"),
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
		URL:          mmSiteURL("/rank/recom/datatype2/"),
		ColumnParser: RTMonthlyRecommendation.ColumnParser,
		Site:         "mm",
	}
	RTTotalRecommendation = RankType{
		URL: mainSiteURL("/rank/recom/datatype3/"),
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
		URL:          mmSiteURL("/rank/recom/datatype3/"),
		ColumnParser: RTTotalRecommendation.ColumnParser,
		Site:         "mm",
	}
	RTTotalBookmark = RankType{
		URL: mainSiteURL("/rank/collect/"),
	}
	RTTotalBookmarkMM = RankType{
		URL:  mmSiteURL("/rank/collect/"),
		Site: "mm",
	}
	RTSignedAuthorNewBook = RankType{
		URL: mainSiteURL("/rank/signnewbook/"),
	}
	RTSignedAuthorNewBookMM = RankType{
		URL:  mmSiteURL("/rank/signnewbook/"),
		Site: "mm",
	}
	RTPublicAuthorNewBook = RankType{
		URL: mainSiteURL("/rank/pubnewbook/"),
	}
	RTPublicAuthorNewBookMM = RankType{
		URL:  mmSiteURL("/rank/pubnewbook/"),
		Site: "mm",
	}
	RTNewSignedAuthorNewBook = RankType{
		URL: mainSiteURL("/rank/newsign/"),
	}
	RTNewSignedAuthorNewBookMM = RankType{
		URL:  mmSiteURL("/rank/newsign/"),
		Site: "mm",
	}
	RTNewAuthorNewBook = RankType{
		URL: mainSiteURL("/rank/newauthor/"),
	}
	RTNewAuthorNewBookMM = RankType{
		URL:  mmSiteURL("/rank/newauthor/"),
		Site: "mm",
	}
	RTWeeklyFans = RankType{
		URL: mainSiteURL("/rank/newfans/"),
	}
	RTWeeklyFansMM = RankType{
		URL:  mmSiteURL("/rank/newfans/"),
		Site: "mm",
	}
	RTLastUpdatedVIP = RankType{
		URL: mainSiteURL("/rank/vipup/"),
	}
	RTDailyMostUpdateVIPMM = RankType{
		URL:  mmSiteURL("/rank/vipup/"),
		Site: "mm",
	}
	RTWeeklyMostUpdateVIPMM = RankType{
		URL:  mmSiteURL("/rank/vipup/datatype2/"),
		Site: "mm",
	}
	RTMonthlyMostUpdateVIPMM = RankType{
		URL:  mmSiteURL("/rank/vipup/datatype3/"),
		Site: "mm",
	}
	RTTotalWordCountMM = RankType{
		URL:  mmSiteURL("/rank/wordcount/"),
		Site: "mm",
	}
	RTTotalBookmarkVIP = RankType{
		URL: mainSiteURL("/rank/vipcollect/"),
	}
	RTWeeklySingleChapterSalesMM = RankType{
		URL:  mmSiteURL("/rank/subscr/"),
		Site: "mm",
	}
	RTTotalSingleChapterSalesVIPMM = RankType{
		URL:  mmSiteURL("/rank/vipsub/"),
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
