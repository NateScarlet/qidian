package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type category struct {
	id   string
	name string
	site string
}

func findCategories() (categories []category, err error) {
	var fetch = func(url, site string) (err error) {
		resp, err := http.Get(url)
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
			categories = append(categories, category{id, name, site})
		})
		return
	}
	err = fetch("https://www.qidian.com/mm/all", "mm")
	if err != nil {
		return
	}
	err = fetch("https://www.qidian.com/all", "")
	if err != nil {
		return
	}
	return
}

type subCategory struct {
	parentID string
	id       string
	name     string
	site     string
}

func findSubCategories(id string) (subCategories []subCategory, err error) {
	var fetch = func(template, site string) (err error) {
		req, err := http.NewRequest("GET", template+id, nil)
		if err != nil {
			return
		}
		req.AddCookie(&http.Cookie{Name: "listStyle", Value: "2"})
		resp, err := http.DefaultClient.Do(req)
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
			var sc = subCategory{name: s.Text(), site: site}
			var pathParts = strings.Split(strings.TrimSuffix(u.Path, "/"), "/")
			var filters = strings.Split(pathParts[len(pathParts)-1], "-")
			for _, filter := range filters {
				if strings.HasPrefix(filter, "chanId") {
					sc.parentID = filter[6:]
				}
				if strings.HasPrefix(filter, "subCateId") {
					sc.id = filter[9:]
				}
			}
			subCategories = append(subCategories, sc)
		})
		return
	}
	err = fetch("https://www.qidian.com/mm/all/chanId", "mm")
	if err != nil {
		return
	}
	err = fetch("https://www.qidian.com/all/chanId", "")
	if err != nil {
		return
	}
	return
}

type dict = map[string]interface{}
type jsonCategory struct {
	Name string `bson:"name" json:"name"`
	Site string `bson:"site" json:"site"`
}

type jsonSubcategory struct {
	ParentID string `bson:"parentID" json:"parentID"`
	Name     string `bson:"name" json:"name"`
	Site     string `bson:"site" json:"site"`
}
type jsonData struct {
	MainCategories dict `bson:"mainCategories" json:"mainCategories"`
	SubCategories  dict `bson:"subCategories" json:"subCategories"`
}

func main() {
	var output string
	flag.StringVar(&output, "o", "", "output")
	flag.Parse()

	var c = &jsonData{
		MainCategories: dict{},
		SubCategories:  dict{},
	}
	cate, err := findCategories()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range cate {
		c.MainCategories[i.id] = jsonCategory{
			Name: i.name,
			Site: i.site,
		}
		subCate, err := findSubCategories(i.id)
		if err != nil {
			log.Fatal(err)
		}
		for _, j := range subCate {
			c.SubCategories[j.id] = jsonSubcategory{
				ParentID: j.parentID,
				Name:     j.name,
				Site:     j.site,
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
