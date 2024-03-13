package database

import (
	"admin/internal/app/payload"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Auth struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewAuth(pool *pgxpool.Pool, ctx context.Context) Auth {
	return Auth{
		pool: pool,
		ctx:  ctx,
	}
}

func (a Auth) Login(auth payload.Auth) bool {
	sql := `
	SELECT
		id
	FROM
	    users 
	WHERE 
	    username = @username
	AND
		password = @password
	`

	args := pgx.NamedArgs{
		"username": auth.Username,
		"password": auth.Password,
	}

	var id int64
	if err := a.pool.QueryRow(a.ctx, sql, args).Scan(&id); err != nil {
		return false
	}

	if id != 0 {
		return true
	}

	return false
}
