package book

// Book model
type Book struct {
	ID                  string
	Title               string
	Author              string
	Description         string
	Category            Category
	SubCategory         SubCategory
	Tags                []string
	CharCount           uint64
	MonthTicketCount    uint64
	WeekRecommendCount  uint64
	TotalRecommendCount uint64
}
