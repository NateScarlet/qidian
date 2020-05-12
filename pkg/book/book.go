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
	CharCount           int64
	MonthTicketCount    int64
	WeekRecommendCount  int64
	TotalRecommendCount int64
}
