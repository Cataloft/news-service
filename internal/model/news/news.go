package news

//go:generate reform

// News represents a row in news table.
//
//reform:news
type News struct {
	ID      int32   `reform:"id,pk"`
	Title   string  `reform:"title"`
	Content *string `reform:"content"`
}
