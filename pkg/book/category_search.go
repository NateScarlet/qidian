package book

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/NateScarlet/qidian/pkg/client"
	"github.com/PuerkitoBio/goquery"
)

// Sort for search
type Sort string

// Sort for search
const (
	SortTotalRecommend Sort = "2"
	SortCharCount      Sort = "3"
	SortWeekRecommend  Sort = "9"
	SortMonthRecommend Sort = "10"
	SortTotalBookmark  Sort = "11"
)

// State for book
type State string

// State for book
const (
	StateAll      State = ""
	StateOnGoing  State = "1"
	StateFinished State = "2"
)

// Sign for book
type Sign string

// Sign for book
const (
	// 全部作品
	SignAll Sign = ""
	// 签约作品
	SignSigned Sign = "1"
	// 精品小说
	SignChoicest Sign = "2"
)

// VIP state for book
type VIP string

// VIP state for book
const (
	VIPAll   VIP = ""
	VIPFalse VIP = "1"
	VIPTrue  VIP = "2"
)

// Update for book
type Update string

// Update for book
const (
	UpdateAll         Update = ""
	UpdateIn3Day      Update = "1"
	UpdateIn7Day      Update = "2"
	UpdateInHalfMonth Update = "3"
	UpdateInMonth     Update = "4"
)

// Size for book
type Size string

// Size for book
const (
	SizeAll          Size = ""
	SizeLt300k       Size = "1"
	SizeGt300kLt500k Size = "2"
	SizeGt500kLt1m   Size = "3"
	SizeGt1mLt2m     Size = "4"
	SizeGt2m         Size = "5"
)

type CategorySearchOptions struct {
	site        string
	sort        Sort
	page        int
	category    Category
	subCategory SubCategory
	state       State
	tag         string
	sign        Sign
	update      Update
	vip         VIP
	size        Size
}

type CategorySearchOption = func(o *CategorySearchOptions)

func CategorySearchURL(opts ...CategorySearchOption) (ret url.URL) {
	var opt = new(CategorySearchOptions)
	for _, i := range opts {
		i(opt)
	}

	ret = mainSiteURL("/all")

	if opt.site != "" {
		ret.Path = opt.site + "/" + ret.Path
	}
	if !strings.HasSuffix(ret.Path, "/") {
		ret.Path += "/"
	}
	var filters = []string{}
	if opt.category != "" {
		filters = append(filters, fmt.Sprintf("chanId%s", string(opt.category)))
	}
	if opt.subCategory != "" {
		filters = append(filters, fmt.Sprintf("subCateId%s", string(opt.subCategory)))
	}
	if opt.state != "" {
		filters = append(filters, fmt.Sprintf("action%s", string(opt.state)))
	}
	if opt.vip != "" {
		filters = append(filters, fmt.Sprintf("vip%s", string(opt.vip)))
	}
	if opt.size != "" {
		filters = append(filters, fmt.Sprintf("size%s", string(opt.size)))
	}
	if opt.sign != "" {
		filters = append(filters, fmt.Sprintf("sign%s", string(opt.sign)))
	}
	if opt.update != "" {
		filters = append(filters, fmt.Sprintf("update%s", string(opt.update)))
	}
	if opt.sort != "" {
		filters = append(filters, fmt.Sprintf("orderId%s", string(opt.sort)))
	}
	if opt.tag != "" {
		filters = append(filters, fmt.Sprintf("tag%s", string(opt.tag)))
	}
	if opt.page > 1 {
		filters = append(filters, fmt.Sprintf("page%d", opt.page))
	}
	if len(filters) > 0 {
		ret.Path += strings.Join(filters, "-") + "/"
	}
	return
}

func CategorySearchOptionPage(v int) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.page = v
	}
}

func CategorySearchOptionSort(v Sort) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.sort = v
	}
}

func CategorySearchOptionCategory(v Category) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.category = v
		o.site = v.Site()
	}
}

func CategorySearchOptionSubCategory(v SubCategory) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.subCategory = v
		CategorySearchOptionCategory(v.Parent())(o)
	}
}

func CategorySearchOptionState(v State) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.state = v
	}
}

func CategorySearchOptionSign(v Sign) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.sign = v
	}
}

func CategorySearchOptionUpdate(v Update) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.update = v
	}
}

func CategorySearchOptionVIP(v VIP) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.vip = v
	}
}

func CategorySearchOptionSize(v Size) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.size = v
	}
}

func CategorySearchOptionTag(v string) CategorySearchOption {
	return func(o *CategorySearchOptions) {
		o.tag = v
	}
}

type CategorySearchResult struct {
	getHTMLResult
	site string
}

// CategorySearch use https://www.qidian.com/all page
func CategorySearch(ctx context.Context, opts ...CategorySearchOption) (ret CategorySearchResult, err error) {
	u := CategorySearchURL(opts...)
	var opt = new(CategorySearchOptions)
	for _, i := range opts {
		i(opt)
	}
	ret.site = opt.site
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

func (r CategorySearchResult) Books() (ret []Book, err error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(r.Body()))
	if err != nil {
		return
	}
	table := doc.Find("table.rank-table-list")
	if table.Length() == 0 {
		return nil, fmt.Errorf("qidian: can not found result table: %s", r.Request().URL)
	}
	return parseTable(table, nil, r.site)
}

func (obj CategorySearchResult) Site() string {
	return obj.site
}
