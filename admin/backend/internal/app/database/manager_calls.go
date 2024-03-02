package database

import (
	"admin/internal/app/payload"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ManagerCalls struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func (c ManagerCalls) GetAll() (calls []payload.ManagerCall, err error) {
	sql := `
	SELECT
	    manager_call.id,
		manager_call.name,
		manager_call.number,
		city_homes.name
	FROM
		manager_call
	JOIN city_homes on manager_call.homeid = city_homes.id
	`

	rows, err := c.pool.Query(c.ctx, sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var call payload.ManagerCall
		err := rows.Scan(&call.Id, &call.Name, &call.Number, &call.HomeName)
		if err != nil {
			continue
		}

		calls = append(calls, call)
	}

	return
}

func NewManagerCalls(ctx context.Context, pool *pgxpool.Pool) ManagerCalls {
	return ManagerCalls{pool: pool, ctx: ctx}
}
