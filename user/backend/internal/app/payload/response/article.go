package response

type Article struct {
	Id      int    `json:id`
	Title   string `json:name`
	Content string `json:content`
}

type ArticleCategory struct {
	Name     string    `json:name`
	Articles []Article `json:articles`
}

type ArticleList struct {
	List []ArticleCategory `json:list`
}
