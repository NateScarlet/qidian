package font

import (
	"io/ioutil"
	"net/http"

	"github.com/golang/groupcache/lru"
	"github.com/golang/groupcache/singleflight"
	"golang.org/x/image/font/sfnt"
)

var fontCache = lru.New(64)

var fontFlight = singleflight.Group{}

// URL for ttf font.
func URL(id string) string {
	return "https://qidian.gtimg.com/qd_anti_spider/" + id + ".ttf"
}

// Get font from ttf font url with a lru cache.
func Get(url string) (*sfnt.Font, error) {
	v, err := fontFlight.Do(url, func() (interface{}, error) {
		if v, ok := fontCache.Get(url); ok {
			return v, nil
		}
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		v, err := sfnt.Parse(data)
		if err != nil {
			return nil, err
		}
		fontCache.Add(url, v)
		return v, nil
	})

	return v.(*sfnt.Font), err
}
