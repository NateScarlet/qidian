package book

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/NateScarlet/qidian/pkg/font"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/net/html"
)

// ParseCount that may contains "万" or thousand period.
func ParseCount(v string) (uint64, error) {
	if v == "- -" {
		return 0, nil
	}
	v = strings.TrimSpace(v)
	is10K := strings.HasSuffix(v, "万")
	if is10K {
		v = v[:len(v)-len("万")]
	}
	v = strings.Replace(v, ",", "", -1)
	ret, err := strconv.ParseFloat(v, 64)
	if is10K {
		ret = ret * 10e3
	}
	return uint64(ret), err
}

var fontPattern = regexp.MustCompile(`url\('([\w.:\/]+\.ttf)'\)`)

func deobfuscate(doc *goquery.Selection) (ret string, err error) {
	doc = doc.Clone()
	styleElem := doc.Find("style").Remove()
	if styleElem.Length() == 0 {
		return doc.Text(), nil
	}
	style, err := styleElem.Html()
	if err != nil {
		return
	}
	style = html.UnescapeString(style)
	match := fontPattern.FindStringSubmatch(style)
	if match == nil {
		h, _ := doc.Html()
		err = fmt.Errorf("qidian: can not found font url: %s", h)
		return
	}
	var f *sfnt.Font
	f, err = font.Get(match[1])
	if err != nil {
		return
	}
	ret, err = font.Deobfuscate(doc.Text(), f)
	return
}

// ParseSelectionCount that may obfuscated by font.
func ParseSelectionCount(doc *goquery.Selection) (ret uint64, err error) {
	var s string
	s, err = deobfuscate(doc)
	if err != nil {
		return
	}
	ret, err = ParseCount(s)
	return
}

// TZ is timezone used to parse timezone, defaults to china standard time.
var TZ = time.FixedZone("CST", 8*60*60)

func parseTimeAt(s string, t time.Time) (ret time.Time, err error) {
	switch {
	case s == "刚刚":
		ret = t
	case strings.HasPrefix(s, "昨日"):
		t.AddDate(0, 0, -1)
		var v time.Time
		v, err = time.Parse("昨日15:04", s)
		if err != nil {
			return
		}
		ret = time.Date(t.Year(), t.Month(), t.Day()-1, v.Hour(), v.Minute(), 0, 0, t.Location())
	case strings.HasSuffix(s, "分钟前"):
		var v int
		v, err = strconv.Atoi(s[:len(s)-len("分钟前")])
		if err != nil {
			return
		}
		ret = t.Add(time.Duration(-v) * time.Minute)
	case strings.HasSuffix(s, "小时前"):
		var v int
		v, err = strconv.Atoi(s[:len(s)-len("小时前")])
		if err != nil {
			return
		}
		ret = t.Add(time.Duration(-v) * time.Hour)
	case len(s) == 10:
		ret, err = time.ParseInLocation("2006-01-02", s, TZ)
	case len(s) == 11:
		ret, err = time.ParseInLocation("01-02 15:04", s, TZ)
		ret = time.Date(
			t.Year(),
			ret.Month(),
			ret.Day(),
			ret.Hour(),
			ret.Minute(),
			0,
			0,
			t.Location(),
		)
	case len(s) == 19:
		ret, err = time.ParseInLocation("2006-01-02 15:04:05", s, TZ)
	default:
		ret, err = time.ParseInLocation("2006-01-02 15:04", s, TZ)
	}
	return
}

// ParseTime for any format that qidian used on their web page.
func ParseTime(s string) (time.Time, error) {
	return parseTimeAt(s, time.Now().In(TZ))
}

// nodeText convert node to multiline text, convert <br> <p> <div> to \n.
func nodeText(n *html.Node) (ret string) {
	switch n.Type {
	case html.TextNode:
		ret += strings.TrimSpace(n.Data)
	case html.ElementNode:
		switch n.Data {
		case "br":
			ret += "\n"
		case "p":
			fallthrough
		case "div":
			ret += "\n"
			fallthrough
		default:
			var cur = n.FirstChild
			for cur != nil {
				ret += nodeText(cur)
				cur = cur.NextSibling
			}
		}
	}
	return
}

func nodesText(s []*html.Node) (ret string) {
	for _, n := range s {
		ret += nodeText(n)
	}
	return strings.TrimSpace(ret)
}

// ColumnParser parse data type from column title.
type ColumnParser interface {
	// Parse title to a column data type
	ParseColumn(book *Book, index int, th *goquery.Selection, td *goquery.Selection) error
}

// ColumnParserFunc implements ColumnParser from a function.
type ColumnParserFunc func(book *Book, index int, th *goquery.Selection, td *goquery.Selection) error

// ParseColumn implements ColumnParser
func (f ColumnParserFunc) ParseColumn(book *Book, index int, th *goquery.Selection, td *goquery.Selection) error {
	return f(book, index, th, td)
}

func defaultColumnParser(book *Book, index int, th *goquery.Selection, s *goquery.Selection) (err error) {
	switch column := strings.TrimSpace(th.Text()); column {
	case "类别":
		parts := strings.SplitN(strings.Trim(s.Text(), "「」"), "·", 2)
		book.Category = CategoryByName(parts[0], book.Site)
		if len(parts) == 2 {
			book.SubCategory = SubCategoryByName(parts[1], book.Site)
		}
	case "小说书名":
		book.Title = s.Text()
		book.ID, _ = s.Find("a").Attr("data-bid")
	case "小说作者":
		a := s.Find("a")
		book.Author.Name = a.Text()
		if href := a.AttrOr("href", ""); strings.HasPrefix(href, "//my.qidian.com/author/") {
			book.Author.ID = href[23:]
		}
	case "字数":
		book.WordCount, err = ParseSelectionCount(s)
		if err != nil {
			return
		}
	case "收藏":
		fallthrough
	case "总收藏":
		book.BookmarkCount, err = ParseSelectionCount(s)
		if err != nil {
			return
		}
	case "推荐":
		var n uint64
		n, err = ParseSelectionCount(s)
		if err != nil {
			return
		}
		if s.HasClass("month") {
			book.MonthRecommendCount = n
		} else if s.HasClass("week") {
			book.WeekRecommendCount = n
		}
	case "周推荐":
		book.WeekRecommendCount, err = ParseSelectionCount(s)
		if err != nil {
			return
		}
	case "月推荐":
		book.MonthRecommendCount, err = ParseSelectionCount(s)
		if err != nil {
			return
		}
	case "总推荐":
		book.TotalRecommendCount, err = ParseSelectionCount(s)
		if err != nil {
			return
		}
	case "更新时间":
		book.LastUpdated, err = ParseTime(s.Text())
		if err != nil {
			return
		}
	case "完本时间":
		book.Finished, err = ParseTime(s.Text())
		if err != nil {
			return
		}
	case "最新章节":
		// TODO: WIP
	case "操作":
	case "排名":
	case "日更字数":
	case "":
		// empty column may appear but contains nothing.
	default:
		return fmt.Errorf("book: unknown column: %s", column)
	}
	return
}

// DefaultColumnParser handler common column
var DefaultColumnParser = ColumnParserFunc(defaultColumnParser)

func parseTable(table *goquery.Selection, columnParser ColumnParser, site string) (ret []Book, err error) {
	var th = make([]*goquery.Selection, 0)
	table.Find("thead > tr > th").Each(func(i int, s *goquery.Selection) {
		th = append(th, s)
	})
	if columnParser == nil {
		columnParser = DefaultColumnParser
	}

	ret = make([]Book, 0, 50)
	table.
		Find("tbody > tr").
		EachWithBreak(func(i int, s *goquery.Selection) bool {
			var book = new(Book)
			book.Site = site
			s.
				ChildrenFiltered("td").
				EachWithBreak(func(i int, s *goquery.Selection) bool {
					if i >= len(th) {
						return false
					}
					err = columnParser.ParseColumn(book, i, th[i], s)
					if err != nil {
						return false
					}
					return true
				})
			if err != nil {
				return false
			}
			ret = append(ret, *book)
			return true
		})

	return
}
