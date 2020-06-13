package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type category struct {
	id   string
	name string
}

func findCategories() (categories []category, err error) {
	resp, err := http.Get("https://www.qidian.com/all")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	doc.Find(`ul[type="category"] > li[data-id]`).Each(func(_ int, s *goquery.Selection) {
		id, _ := s.Attr("data-id")
		name := s.Find("a").Text()
		categories = append(categories, category{id, name})
	})
	return
}

type subCategory struct {
	parentID string
	id       string
	name     string
}

func findSubCategories(id string) (subCategories []subCategory, err error) {
	resp, err := http.Get("https://www.qidian.com/all?chanId=" + id)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return
	}
	doc.Find(`.sub-type > dl:not(.hidden) > dd[data-subtype] > a`).Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		u, _ := url.Parse(href)
		q := u.Query()
		subCategories = append(subCategories, subCategory{
			parentID: q.Get("chanId"),
			id:       q.Get("subCateId"),
			name:     s.Text(),
		})
	})

	return
}

type dict = map[string]interface{}
type templateContextSubCategory struct {
	ParentID string
	Name     string
}
type templateContext struct {
	MainCategories dict
	SubCategories  dict
}

func main() {
	var output string
	flag.StringVar(&output, "o", "", "output")
	flag.Parse()

	var c = &templateContext{
		MainCategories: dict{},
		SubCategories:  dict{},
	}
	cate, err := findCategories()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range cate {
		c.MainCategories[i.id] = i.name
		subCate, err := findSubCategories(i.id)
		if err != nil {
			log.Fatal(err)
		}
		for _, j := range subCate {
			c.SubCategories[j.id] = templateContextSubCategory{
				ParentID: j.parentID,
				Name:     j.name,
			}
		}
	}

	var w io.Writer = os.Stdout
	if output != "" {
		var f *os.File
		f, err = os.OpenFile(output, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	err = enc.Encode(c)
	if err != nil {
		log.Fatal(err)
	}
}
