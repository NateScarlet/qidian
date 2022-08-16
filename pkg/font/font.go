package font

import (
	"context"

	"github.com/NateScarlet/qidian/pkg/client"
	"golang.org/x/image/font/sfnt"
)

// URL for ttf font.
func URL(id string) string {
	return "https://qidian.gtimg.com/qd_anti_spider/" + id + ".ttf"
}

func Get(ctx context.Context, url string) (_ *sfnt.Font, err error) {
	data, err := client.GetAsset(ctx, url)
	if err != nil {
		return
	}
	return sfnt.Parse(data)
}
