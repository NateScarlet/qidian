package main

import (
	"fmt"
	"net/http"
	"net/url"

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

func main() {
	cate, err := findCategories()
	if err != nil {
		panic(err)
	}
	for _, i := range cate {
		fmt.Printf("%s \t%s\n", i.id, i.name)
		subCate, _ := findSubCategories(i.id)
		for _, j := range subCate {
			fmt.Printf("%s \t%s \t%s\n", j.parentID, j.id, j.name)
		}
	}
}
