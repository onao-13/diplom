package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Article struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewArticle(pool *pgxpool.Pool, ctx context.Context) Article {
	return Article{pool: pool, ctx: ctx}
}

//func (a *Article) Get(id int64) (payload.Article, error) {
//	var article payload.Article
//
//	sql := `SELECT * FROM articles WHERE id=$1`
//
//	if err := pgxscan.Get(a.ctx, a.pool, &article, sql, id); err != nil {
//		return payload.Article{}, err
//	}
//
//	return article, nil
//}
//
//func (a *Article) GetAllCategories() ([]entity.ArticleCategory, error) {
//	var categories []entity.ArticleCategory
//
//	sql := `SELECT * FROM articles_categories`
//
//	if err := pgxscan.Select(a.ctx, a.pool, &categories, sql); err != nil {
//		return nil, err
//	}
//
//	return categories, nil
//}
//
//func (a *Article) GetLimitArticlesByCategoryId(categoryId int64, limit int64) ([]payload.Article, error) {
//	var articles []payload.Article
//
//	sql := `SELECT * FROM articles WHERE categoryId = $1 LIMIT $2`
//
//	if err := pgxscan.Select(a.ctx, a.pool, &articles, sql, categoryId, limit); err != nil {
//		return nil, err
//	}
//
//	return articles, nil
//}
