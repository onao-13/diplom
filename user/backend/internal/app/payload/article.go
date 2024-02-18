package payload

type (
	Article struct {
		Id      int    `json:"id"`
		Title   string `json:"name"`
		Content string `json:"content"`
	}

	ArticleCategory struct {
		Name     string    `json:"name"`
		Articles []Article `json:"articles"`
	}

	ArticleList struct {
		List []ArticleCategory `json:"list"`
	}
)
