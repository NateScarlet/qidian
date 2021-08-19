package book

import (
	"strings"
	"testing"

	"github.com/NateScarlet/snapshot/pkg/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func categorySearchID(opts []CategorySearchOption) string {
	u := CategorySearchURL(opts...)
	u.Scheme = ""
	u.Host = ""
	return strings.Replace(u.String(), "/", "__", -1)
}

func TestCategorySearch_simple(t *testing.T) {
	for _, c := range [][]CategorySearchOption{
		{},
		{CategorySearchOptionCategory(C科幻)},
		{CategorySearchOptionSubCategory(SC未来世界)},
		{CategorySearchOptionPage(2)},
		{CategorySearchOptionPage(2), CategorySearchOptionSort(SortMonthRecommend)},
		{CategorySearchOptionSort(SortLastUpdated)},
		{CategorySearchOptionSort(SortMonthRecommend)},
		{CategorySearchOptionSort(SortRecentFinished)},
		{CategorySearchOptionSort(SortTotalBookmark)},
		{CategorySearchOptionSort(SortTotalRecommend)},
		{CategorySearchOptionSort(SortWeekRecommend)},
		{CategorySearchOptionSign(SignSigned)},
		{CategorySearchOptionSign(SignChoicest)},
		{CategorySearchOptionUpdate(UpdateIn3Day)},
		{CategorySearchOptionUpdate(UpdateIn7Day)},
		{CategorySearchOptionUpdate(UpdateInHalfMonth)},
		{CategorySearchOptionUpdate(UpdateInMonth)},
		{CategorySearchOptionState(StateOnGoing)},
		{CategorySearchOptionState(StateFinished)},
		{CategorySearchOptionSize(SizeLt300k)},
		{CategorySearchOptionSize(SizeGt300kLt500k)},
		{CategorySearchOptionSize(SizeGt500kLt1m)},
		{CategorySearchOptionSize(SizeGt1mLt2m)},
		{CategorySearchOptionSize(SizeGt2m)},
		{CategorySearchOptionVIP(VIPFalse)},
		{CategorySearchOptionVIP(VIPTrue)},
		{CategorySearchOptionTag("变身")},
	} {
		t.Run(categorySearchID(c), func(t *testing.T) {
			res, err := CategorySearch(context.Background(), c...)
			require.NoError(t, err)
			books, err := res.Books()
			require.NoError(t, err)
			var opt = new(CategorySearchOptions)
			for _, i := range c {
				i(opt)
			}
			assert.Len(t, books, 20)
			for _, i := range books {
				assert.NotEmpty(t, i.ID)
				assert.NotEmpty(t, i.Title)
				assert.NotEmpty(t, i.Author.Name)
				assert.NotEmpty(t, i.Author.ID)
				assert.NotEmpty(t, i.Category)
				assert.NotEmpty(t, i.SubCategory)
				assert.NotEmpty(t, i.WordCount)
				if opt.sort == SortTotalBookmark {
					assert.NotEmpty(t, i.BookmarkCount)
				}
				if opt.sort == "" {
					assert.NotEmpty(t, i.LastUpdated)
				}
				if opt.sort == SortWeekRecommend {
					assert.NotEmpty(t, i.WeekRecommendCount)
				}
				if opt.sort == SortMonthRecommend {
					assert.NotEmpty(t, i.MonthRecommendCount)
				}
				if opt.sort == SortTotalRecommend {
					assert.NotEmpty(t, i.TotalRecommendCount)
				}
				if opt.sort == SortRecentFinished {
					assert.NotEmpty(t, i.Finished)
				}
			}
			if snapshot.DefaultUpdate {
				snapshot.MatchJSON(t, books)
			}
		})

	}

}
