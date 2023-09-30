package database

import (
	// "backend/internal/app/payload/response"
	"context"

	// "github.com/georgysavva/scany/v2/pgxscan"
	// "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Article struct {
	pool *pgxpool.Pool
	log logrus.Logger
	ctx context.Context
}

func NewArticle(pool *pgxpool.Pool, log logrus.Logger, ctx context.Context) Article {
	return Article{pool: pool, log: log, ctx: ctx}
}

// func (a *Article) View(id int) (response.Article, error) {
// 	var article response.Article

// 	sql := `SELECT * FROM articles WHERE id=$1`

// 	if err := pgxscan.Select(a.ctx, a.pool, &article, sql, id); err != nil {

// 	}
// }