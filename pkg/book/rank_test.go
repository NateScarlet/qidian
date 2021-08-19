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
		rank                           Rank
		shouldHasMonthlyTicket         bool
		shouldHasWeeklyRecommendation  bool
		shouldHasMonthlyRecommendation bool
		shouldHasTotalRecommendation   bool
		shouldHasBookmark              bool
	}{
		{
			name: "monthly-ticket",
			rank: Rank{
				Type: RTMonthlyTicket,
			},
			shouldHasMonthlyTicket: true,
		},
		{
			name: "monthly-ticket-mm",
			rank: Rank{
				Type: RTMonthlyTicketMM,
			},
			shouldHasMonthlyTicket: true,
		},
		{
			name: "monthly-ticket-vip",
			rank: Rank{
				Type: RTMonthlyTicketVIP,
			},
			shouldHasMonthlyTicket: true,
		},
		{
			name: "monthly-ticket category",
			rank: Rank{
				Type:     RTMonthlyTicket,
				Category: C科幻,
			},
			shouldHasMonthlyTicket: true,
		},
		{
			name: "monthly-ticket history",
			rank: Rank{
				Type:  RTMonthlyTicket,
				Year:  2020,
				Month: time.January,
			},
			shouldHasMonthlyTicket: true,
		},
		{
			name: "new-book-sales-mm",
			rank: Rank{
				Type: RTDailySales,
			},
		},
		{
			name: "daily-sales",
			rank: Rank{
				Type: RTDailySales,
			},
		},
		{
			name: "daily-sales-mm",
			rank: Rank{
				Type: RTDailySales,
			},
		},
		{
			name: "weekly-read",
			rank: Rank{
				Type: RTWeeklyRead,
			},
		},
		{
			name: "weekly-read-mm",
			rank: Rank{
				Type: RTWeeklyReadMM,
			},
		},
		{
			name: "weekly-recommendation",
			rank: Rank{
				Type: RTWeeklyRecommendation,
			},
			shouldHasWeeklyRecommendation: true,
		},
		{
			name: "weekly-recommendation-mm",
			rank: Rank{
				Type: RTWeeklyRecommendationMM,
			},
			shouldHasWeeklyRecommendation: true,
		},
		{
			name: "monthly-recommendation",
			rank: Rank{
				Type: RTMonthlyRecommendation,
			},
			shouldHasMonthlyRecommendation: true,
		},
		{
			name: "monthly-recommendation-mm",
			rank: Rank{
				Type: RTMonthlyRecommendationMM,
			},
			shouldHasMonthlyRecommendation: true,
		},
		{
			name: "total-recommendation",
			rank: Rank{
				Type: RTTotalRecommendation,
			},
			shouldHasTotalRecommendation: true,
		},
		{
			name: "total-recommendation-mm",
			rank: Rank{
				Type: RTTotalRecommendationMM,
			},
			shouldHasTotalRecommendation: true,
		},
		{
			name: "signed-author-new-book",
			rank: Rank{
				Type: RTSignedAuthorNewBook,
			},
		},
		{
			name: "signed-author-new-book-mm",
			rank: Rank{
				Type: RTSignedAuthorNewBookMM,
			},
		},
		{
			name: "public-author-new-book",
			rank: Rank{
				Type: RTPublicAuthorNewBook,
			},
		},
		{
			name: "public-author-new-book-mm",
			rank: Rank{
				Type: RTPublicAuthorNewBookMM,
			},
		},
		{
			name: "new-author-new-book",
			rank: Rank{
				Type: RTNewAuthorNewBook,
			},
		},
		{
			name: "new-author-new-book-mm",
			rank: Rank{
				Type: RTNewAuthorNewBookMM,
			},
		},
		{
			name: "new-signed-author-new-book",
			rank: Rank{
				Type: RTNewSignedAuthorNewBook,
			},
		},
		{
			name: "new-signed-author-new-book-mm",
			rank: Rank{
				Type: RTNewSignedAuthorNewBookMM,
			},
		},
		{
			name: "weekly-fans",
			rank: Rank{
				Type: RTWeeklyFans,
			},
		},
		{
			name: "weekly-fans-mm",
			rank: Rank{
				Type: RTWeeklyFansMM,
			},
		},
		{
			name: "last-updated",
			rank: Rank{
				Type: RTLastUpdatedVIP,
			},
		},
		{
			name: "total-bookmark-vip",
			rank: Rank{
				Type: RTTotalBookmarkVIP,
			},
		},
		{
			name: "weekly-reward-vip",
			rank: Rank{
				Type: RTWeeklyRewardVIP,
			},
		},
		{
			name: "weekly-single-chapter-sales-mm",
			rank: Rank{
				Type: RTWeeklySingleChapterSalesMM,
			},
		},
		{
			name: "total-single-chapter-sales-vip-mm",
			rank: Rank{
				Type: RTTotalSingleChapterSalesVIPMM,
			},
		},
		{
			name: "daily-most-update-vip-mm",
			rank: Rank{
				Type: RTDailyMostUpdateVIPMM,
			},
		},
		{
			name: "weekly-most-update-vip-mm",
			rank: Rank{
				Type: RTWeeklyMostUpdateVIPMM,
			},
		},
		{
			name: "monthly-most-update-vip-mm",
			rank: Rank{
				Type: RTMonthlyMostUpdateVIPMM,
			},
		},
		{
			name: "total-word-count-mm",
			rank: Rank{
				Type: RTTotalWordCountMM,
			},
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			books, err := c.rank.Fetch(context.Background())
			require.NoError(t, err)
			assert.NotEmpty(t, books)
			if snapshot.DefaultUpdate {
				snapshot.MatchJSON(t, books)
			}
			for _, book := range books {
				assert.NotEmpty(t, book.ID)
				assert.NotEmpty(t, book.Title)
				assert.NotEmpty(t, book.Category)
				assert.Equal(t, c.rank.Type.Site, book.Site)
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
