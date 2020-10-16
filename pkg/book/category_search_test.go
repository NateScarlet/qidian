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
		NewCategorySearch().SetPage(2).SetSort(SortMonthRecommend),
		NewCategorySearch().SetSort(SortLastUpdated),
		NewCategorySearch().SetSort(SortMonthRecommend),
		NewCategorySearch().SetSort(SortRecentFinished),
		NewCategorySearch().SetSort(SortTotalBookmark),
		NewCategorySearch().SetSort(SortTotalRecommend),
		NewCategorySearch().SetSort(SortWeekRecommend),
		NewCategorySearch().SetSign(SignSigned),
		NewCategorySearch().SetSign(SignChoicest),
		NewCategorySearch().SetUpdate(UpdateIn3Day),
		NewCategorySearch().SetUpdate(UpdateIn7Day),
		NewCategorySearch().SetUpdate(UpdateInHalfMonth),
		NewCategorySearch().SetUpdate(UpdateInMonth),
		NewCategorySearch().SetState(StateOnGoing),
		NewCategorySearch().SetState(StateFinished),
		NewCategorySearch().SetSize(SizeLt300k),
		NewCategorySearch().SetSize(SizeGt300kLt500k),
		NewCategorySearch().SetSize(SizeGt500kLt1m),
		NewCategorySearch().SetSize(SizeGt1mLt2m),
		NewCategorySearch().SetSize(SizeGt2m),
		NewCategorySearch().SetVIP(VIPFalse),
		NewCategorySearch().SetVIP(VIPTrue),
		NewCategorySearch().SetTag("变身"),
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
				if s.Sort == SortTotalBookmark {
					assert.NotEmpty(t, i.BookmarkCount)
				}
				if s.Sort == "" {
					assert.NotEmpty(t, i.LastUpdated)
				}
				if s.Sort == SortWeekRecommend {
					assert.NotEmpty(t, i.WeekRecommendCount)
				}
				if s.Sort == SortMonthRecommend {
					assert.NotEmpty(t, i.MonthRecommendCount)
				}
				if s.Sort == SortTotalRecommend {
					assert.NotEmpty(t, i.TotalRecommendCount)
				}
				if s.Sort == SortRecentFinished {
					assert.NotEmpty(t, i.Finished)
				}
			}
		})

	}

}
