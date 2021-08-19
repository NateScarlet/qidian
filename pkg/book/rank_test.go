package book

import (
	"context"
	"testing"
	"time"

	"github.com/NateScarlet/snapshot/pkg/snapshot"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRank_Fetch(t *testing.T) {
	for _, c := range []struct {
		name                           string
		rankType                       RankType
		options                        []RankOption
		shouldHasMonthlyTicket         bool
		shouldHasWeeklyRecommendation  bool
		shouldHasMonthlyRecommendation bool
		shouldHasTotalRecommendation   bool
		shouldHasBookmark              bool
	}{
		{
			name:                   "monthly-ticket",
			rankType:               RTMonthlyTicket,
			shouldHasMonthlyTicket: true,
		},
		{
			name:                   "monthly-ticket-page2",
			rankType:               RTMonthlyTicket,
			options:                []RankOption{RankOptionPage(2)},
			shouldHasMonthlyTicket: true,
		},
		{
			name:                   "monthly-ticket-mm",
			rankType:               RTMonthlyTicketMM,
			shouldHasMonthlyTicket: true,
		},
		{
			name:                   "monthly-ticket-vip",
			rankType:               RTMonthlyTicketVIP,
			shouldHasMonthlyTicket: true,
		},
		{
			name:                   "monthly-ticket category",
			rankType:               RTMonthlyTicket,
			options:                []RankOption{RankOptionCategory(C科幻)},
			shouldHasMonthlyTicket: true,
		},
		{
			name:                   "monthly-ticket history",
			rankType:               RTMonthlyTicket,
			options:                []RankOption{RankOptionYearMonth(2020, time.January)},
			shouldHasMonthlyTicket: true,
		},
		{
			name:     "new-book-sales-mm",
			rankType: RTDailySales,
		},
		{
			name:     "daily-sales",
			rankType: RTDailySales,
		},
		{
			name:     "daily-sales-mm",
			rankType: RTDailySales,
		},
		{
			name:     "weekly-read",
			rankType: RTWeeklyRead,
		},
		{
			name:     "weekly-read-mm",
			rankType: RTWeeklyReadMM,
		},
		{
			name:                          "weekly-recommendation",
			rankType:                      RTWeeklyRecommendation,
			shouldHasWeeklyRecommendation: true,
		},
		{
			name:                          "weekly-recommendation-mm",
			rankType:                      RTWeeklyRecommendationMM,
			shouldHasWeeklyRecommendation: true,
		},
		{
			name:                           "monthly-recommendation",
			rankType:                       RTMonthlyRecommendation,
			shouldHasMonthlyRecommendation: true,
		},
		{
			name:                           "monthly-recommendation-mm",
			rankType:                       RTMonthlyRecommendationMM,
			shouldHasMonthlyRecommendation: true,
		},
		{
			name:                         "total-recommendation",
			rankType:                     RTTotalRecommendation,
			shouldHasTotalRecommendation: true,
		},
		{
			name:                         "total-recommendation-mm",
			rankType:                     RTTotalRecommendationMM,
			shouldHasTotalRecommendation: true,
		},
		{
			name:     "signed-author-new-book",
			rankType: RTSignedAuthorNewBook,
		},
		{
			name:     "signed-author-new-book-mm",
			rankType: RTSignedAuthorNewBookMM,
		},
		{
			name:     "public-author-new-book",
			rankType: RTPublicAuthorNewBook,
		},
		{
			name:     "public-author-new-book-mm",
			rankType: RTPublicAuthorNewBookMM,
		},
		{
			name:     "new-author-new-book",
			rankType: RTNewAuthorNewBook,
		},
		{
			name:     "new-author-new-book-mm",
			rankType: RTNewAuthorNewBookMM,
		},
		{
			name:     "new-signed-author-new-book",
			rankType: RTNewSignedAuthorNewBook,
		},
		{
			name:     "new-signed-author-new-book-mm",
			rankType: RTNewSignedAuthorNewBookMM,
		},
		{
			name:     "weekly-fans",
			rankType: RTWeeklyFans,
		},
		{
			name:     "weekly-fans-mm",
			rankType: RTWeeklyFansMM,
		},
		{
			name:     "last-updated",
			rankType: RTLastUpdatedVIP,
		},
		{
			name:     "total-bookmark-vip",
			rankType: RTTotalBookmarkVIP,
		},
		{
			name:     "weekly-reward-vip",
			rankType: RTWeeklyRewardVIP,
		},
		{
			name:     "weekly-single-chapter-sales-mm",
			rankType: RTWeeklySingleChapterSalesMM,
		},
		{
			name:     "total-single-chapter-sales-vip-mm",
			rankType: RTTotalSingleChapterSalesVIPMM,
		},
		{
			name:     "daily-most-update-vip-mm",
			rankType: RTDailyMostUpdateVIPMM,
		},
		{
			name:     "weekly-most-update-vip-mm",
			rankType: RTWeeklyMostUpdateVIPMM,
		},
		{
			name:     "monthly-most-update-vip-mm",
			rankType: RTMonthlyMostUpdateVIPMM,
		},
		{
			name:     "total-word-count-mm",
			rankType: RTTotalWordCountMM,
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			res, err := Rank(context.Background(), c.rankType, c.options...)
			require.NoError(t, err)
			books, err := res.Books()
			assert.NotEmpty(t, books)
			if snapshot.DefaultUpdate {
				snapshot.MatchJSON(t, books)
			}
			for _, book := range books {
				assert.NotEmpty(t, book.ID)
				assert.NotEmpty(t, book.Title)
				assert.NotEmpty(t, book.Category)
				assert.Equal(t, c.rankType.Site, book.Site)
				assert.NotEmpty(t, book.LastUpdated)
				if c.shouldHasMonthlyTicket {
					assert.NotEmpty(t, book.MonthTicketCount)
				}
				if c.shouldHasWeeklyRecommendation {
					assert.NotEmpty(t, book.WeekRecommendCount)
				}
				if c.shouldHasMonthlyRecommendation {
					assert.NotEmpty(t, book.MonthRecommendCount)
				}
				if c.shouldHasTotalRecommendation {
					assert.NotEmpty(t, book.TotalRecommendCount)
				}
				if c.shouldHasBookmark {
					assert.NotEmpty(t, book.BookmarkCount)
				}
			}
		})
	}
}
