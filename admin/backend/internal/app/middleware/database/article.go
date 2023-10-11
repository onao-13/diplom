package database

import (
	"admin/internal/payload"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type Article struct {
	ctx context.Context
	pool *pgxpool.Pool
	log logrus.Logger
}

func NewArticle(ctx context.Context, pool *pgxpool.Pool, log logrus.Logger) Article {
	return Article{ctx: ctx, pool: pool, log: log}
}

func (a *Article) Create(art payload.Article) error {
	sql := `INSERT INTO articles(title, content) VALUES (@title, @content)`
	args := pgx.NamedArgs{
		"title": art.Title,
		"content": art.Content,
	}

	if _, err := a.pool.Exec(a.ctx, sql, args); err != nil {
		return err
	}

	return nil
}

func (a *Article) Preview(id int64) (payload.Article, error) {
	var art payload.Article
	
	sql := `SELECT title, content FROM articles WHERE id=$1`
	
	if err := pgxscan.Select(a.ctx, a.pool, &art, sql, id); err != nil {
		return payload.Article{}, err
	}

	return art, nil
}

func (a *Article) Update(id int64, art payload.Article) error {
	sql := `UPDATE articles 
			SET title=@title, content=@content
			WHERE id=@id
			`
	
	args := pgx.NamedArgs{
		"title": art.Title,
		"content": art.Content,
		"id": id,
	}

	if _, err := a.pool.Exec(a.ctx, sql, args); err != nil {
		return err
	} 

	return nil
}

func (a *Article) Delete(id int64) error {
	sql := `DELETE FROM articles WHERE id=@id`

	arg := pgx.NamedArgs{
		"id": id,
	}

	if _, err := a.pool.Exec(a.ctx, sql, arg); err != nil {
		return err
	}

	return nil
}
