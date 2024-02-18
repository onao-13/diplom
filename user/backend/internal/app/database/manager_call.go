package database

import (
	"backend/internal/app/payload"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ManagerCall struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewManagerCall(ctx context.Context, pool *pgxpool.Pool) ManagerCall {
	return ManagerCall{ctx: ctx, pool: pool}
}

func (m ManagerCall) Save(call payload.SendManagerCall) error {
	sql := `
	INSERT INTO
		manager_call(name, number, homeid) 
	VALUES(@name, @number, @homeid)
	`

	args := pgx.NamedArgs{
		"name":   call.Name,
		"number": call.Number,
		"homeid": call.HomeId,
	}

	if _, err := m.pool.Exec(m.ctx, sql, args); err != nil {
		return err
	}

	return nil
}
