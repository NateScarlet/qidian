package book

import "time"

// Book model
type Book struct {
	ID          string
	Title       string
	Author      string
	Description string
	Category    Category
	SubCategory SubCategory
	Tags        []string
	LastUpdated time.Time
	CharCount   uint64
	// only avaliable when search by bookmark
	BookmarkCount       uint64
	MonthTicketCount    uint64
	WeekRecommendCount  uint64
	TotalRecommendCount uint64
}
