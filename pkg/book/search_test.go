package book

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func TestSearch_simple(t *testing.T) {
	for _, c := range []*Search{
		NewSearch(),
		NewSearch().SetCategory(C科幻),
		NewSearch().SetSubCategory(SC未来世界),
		NewSearch().SetPage(2),
		NewSearch().SetPage(2).SetSort(SMonthRecommend),
	} {
		s := c
		t.Run(fmt.Sprintf("%+v", c), func(t *testing.T) {
			res, err := s.Execute(context.Background())
			require.NoError(t, err)
			assert.Len(t, res, 50)
			for _, i := range res {
				assert.NotEmpty(t, i.ID)
				assert.NotEmpty(t, i.Title)
				assert.NotEmpty(t, i.Author)
				assert.NotEmpty(t, i.Category)
				assert.NotEmpty(t, i.SubCategory)
			}
		})

	}

}
