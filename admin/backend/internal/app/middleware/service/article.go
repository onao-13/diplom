package service

import (
	"admin/internal/app/middleware/database"
	"admin/internal/payload"
)

type Article struct{
	db database.Article
}

func NewArticle(db database.Article) Article {
	return Article{db}
}

func (a *Article) Create(art payload.Article) error {
	return a.db.Create(art)
}

func (a *Article) Preview(id int64) (payload.Article, error) {
	return a.db.Preview(id)
}

func (a *Article) Update(id int64, art payload.Article) error {
	return a.db.Update(id, art)
}

func (a *Article) Delete(id int64) error {
	return a.db.Delete(id)
}
