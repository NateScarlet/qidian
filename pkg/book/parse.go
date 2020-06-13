package book

import (
	"errors"
	"html"
	"regexp"
	"strconv"
	"strings"

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
