package newscategories

//go:generate reform

// NewsCategories represents a row in news_categories table.
//
//reform:news_categories
type NewsCategories struct {
	NewsID     int32 `reform:"news_id"`
	CategoryID int32 `reform:"category_id"`
}
