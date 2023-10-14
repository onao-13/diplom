package database

import (
	"backend/internal/app/entity"
	"backend/internal/app/payload/response"
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Home struct {
	ctx context.Context
	pool *pgxpool.Pool
}

func NewHome(ctx context.Context, pool *pgxpool.Pool) Home {
	return Home{ctx, pool}
}

func (h *Home) GetById(id int64) (response.Home, error) {
	var home response.Home

	sql := `SELECT * FROM homes WHERE id=$1`

	if err := pgxscan.Get(h.ctx, h.pool, &home, sql, id); err != nil {
		return response.Home{}, err
	}

	return home, nil
}

func (h *Home) GetAllCity() (entity.City, error) {
	var city entity.City

	sql := `SELECT * FROM cities`

	if err := pgxscan.Select(h.ctx, h.pool, &city, sql); err != nil {
		return entity.City{}, err
	}

	return city, nil
}

func (h *Home) GetHomesByCityName(cityId int64) ([]entity.Home, error) {
	var list []entity.Home

	sql := `SELECT * FROM homes WHERE cityId=$1`

	if err := pgxscan.Select(h.ctx, h.pool, &list, sql, cityId); err != nil {
		return nil, err
	}

	return list, nil
}
