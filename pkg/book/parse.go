package book

import (
	"errors"
	"html"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/NateScarlet/qidian/pkg/font"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/image/font/sfnt"
)

func parseCount(v string) (uint64, error) {
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

func parseCountSelection(doc *goquery.Selection) (ret uint64, err error) {
	doc = doc.Clone()
	styleElem := doc.Find("style")
	style, err := styleElem.Html()
	if err != nil {
		return
	}
	style = html.UnescapeString(style)
	match := fontPattern.FindStringSubmatch(style)
	if match == nil {
		err = errors.New("can not found font url")
		return
	}
	var f *sfnt.Font
	f, err = font.Get(match[1])
	if err != nil {
		return
	}
	var text string
	styleElem.Remove()
	text, err = font.Deobfuscate(doc.Text(), f)
	if err != nil {
		return
	}
	ret, err = parseCount(text)
	return
}

// TZ is timezone used to parse timezone, defaults to Asia/Shanghai.
var TZ, _ = time.LoadLocation("Asia/Shanghai")

func parseTimeAt(s string, t time.Time) (ret time.Time, err error) {
	if strings.HasPrefix(s, "昨日") {
		t.AddDate(0, 0, -1)
		var v time.Time
		v, err = time.Parse("昨日15:04", s)
		if err != nil {
			return
		}
		ret = time.Date(t.Year(), t.Month(), t.Day()-1, v.Hour(), v.Minute(), 0, 0, t.Location())
		return
	} else if strings.HasSuffix(s, "分钟前") {
		var v int
		v, err = strconv.Atoi(s[:len(s)-len("分钟前")])
		if err != nil {
			return
		}
		ret = t.Add(time.Duration(-v) * time.Minute)
		return
	} else if strings.HasSuffix(s, "小时前") {
		var v int
		v, err = strconv.Atoi(s[:len(s)-len("小时前")])
		if err != nil {
			return
		}
		ret = t.Add(time.Duration(-v) * time.Hour)
		return
	} else {
		ret, err = time.ParseInLocation("2006-01-02", s, TZ)
		return
	}
}

func parseTime(s string) (time.Time, error) {
	return parseTimeAt(s, time.Now().In(TZ))
}
