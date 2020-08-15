package book

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func TestCategorySearch_simple(t *testing.T) {
	for _, c := range []*CategorySearch{
		NewCategorySearch(),
		NewCategorySearch().SetCategory(C科幻),
		NewCategorySearch().SetSubCategory(SC未来世界),
		NewCategorySearch().SetPage(2),
		NewCategorySearch().SetPage(2).SetSort(SMonthRecommend),
		NewCategorySearch().SetSort(SLastUpdated),
		NewCategorySearch().SetSort(SMonthRecommend),
		NewCategorySearch().SetSort(SRecentFinished),
		NewCategorySearch().SetSort(STotalBookmark),
		NewCategorySearch().SetSort(STotalRecommend),
		NewCategorySearch().SetSort(SWeekRecommend),
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
				assert.NotEmpty(t, i.WordCount)
				if s.Sort == STotalBookmark {
					assert.NotEmpty(t, i.BookmarkCount)
				}
				if s.Sort == "" {
					assert.NotEmpty(t, i.LastUpdated)
				}
				if s.Sort == SWeekRecommend {
					assert.NotEmpty(t, i.WeekRecommendCount)
				}
				if s.Sort == SMonthRecommend {
					assert.NotEmpty(t, i.MonthRecommendCount)
				}
				if s.Sort == STotalRecommend {
					assert.NotEmpty(t, i.TotalRecommendCount)
				}
				if s.Sort == SRecentFinished {
					assert.NotEmpty(t, i.Finished)
				}
			}
		})

	}

}
