package database

import (
	"backend/internal/app/payload"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Home struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewHome(ctx context.Context, pool *pgxpool.Pool) Home {
	return Home{ctx, pool}
}

func (h Home) GetById(homeId string) (home *payload.Home, err error) {
	sql := `SELECT
		city_homes.id,
		city_homes.name,
		city_homes.street,
		city_homes.price,
		city_homes.cityid,
		COALESCE(city_homes.description, ''),
		COALESCE(city_homes.layout, ''),
		COALESCE(city_homes.greenzone, ''),
		COALESCE(city_homes.infrastructure, ''),
		COALESCE(city_homes.events, ''),
		COALESCE(city_homes.schools, ''),
		COALESCE(city_homes.transports, ''),
		COALESCE(city_homes.popularLocations, '')
	FROM city_homes
	WHERE 
	    id=@id
	`

	args := pgx.NamedArgs{
		"id": homeId,
	}

	rows, err := h.pool.Query(h.ctx, sql, args)
	if err != nil {
		return
	}

	var batch = &pgx.Batch{}

	for rows.Next() {
		home = &payload.Home{}

		err = rows.Scan(
			&home.Id, &home.Name, &home.Street, &home.Price, &home.CityId,
			&home.Description, &home.Layout, &home.GreenZone,
			&home.Infrastructure, &home.Events, &home.Schools,
			&home.Transports, &home.PopularLocations,
		)
		if err != nil {
			fmt.Println("Err scan: ", err)
			return
		}

		h.getImage(batch, home.Id)
	}

	br := h.pool.SendBatch(h.ctx, batch)
	defer br.Close()

	for i := 0; i < batch.Len(); i++ {
		rows, err = br.Query()
		if err != nil {
			return
		}

		for rows.Next() {
			var (
				img    payload.HomeImage
				homeId string
			)
			if err = rows.Scan(&img.Id, &img.URL, &homeId); err != nil {
				return
			}

			home.Images = append(home.Images, img)
		}
	}

	return
}

func (h *Home) GetAll(cityId int64) (homes []*payload.Home, err error) {
	sql := `
	SELECT
		city_homes.id,
		city_homes.name,
		city_homes.street,
		city_homes.price,
		city_homes.cityid,
		COALESCE(city_homes.description, ''),
		COALESCE(city_homes.layout, ''),
		COALESCE(city_homes.greenzone, ''),
		COALESCE(city_homes.infrastructure, ''),
		COALESCE(city_homes.events, ''),
		COALESCE(city_homes.schools, ''),
		COALESCE(city_homes.transports, ''),
		COALESCE(city_homes.popularLocations, '')
	FROM city_homes
	WHERE 
	    cityid=@cityid
	LIMIT 6
	`

	arg := pgx.NamedArgs{
		"cityid": cityId,
	}

	rows, err := h.pool.Query(h.ctx, sql, arg)
	if err != nil {
		return
	}

	var (
		batch    = &pgx.Batch{}
		homesIds = make(map[string]*payload.Home)
	)

	for rows.Next() {
		var home payload.Home

		err = rows.Scan(
			&home.Id, &home.Name, &home.Street, &home.Price,
			&home.CityId, &home.Description, &home.Layout, &home.GreenZone,
			&home.Infrastructure, &home.Events, &home.Schools, &home.Transports,
			&home.PopularLocations,
		)
		if err != nil {
			fmt.Println("Err scan: ", err)
			continue
		}

		homesIds[home.Id] = &home
		h.getImage(batch, home.Id)
	}

	br := h.pool.SendBatch(h.ctx, batch)
	defer br.Close()

	for i := 0; i < batch.Len(); i++ {
		rows, err = br.Query()
		if err != nil {
			return
		}

		for rows.Next() {
			var (
				img    payload.HomeImage
				homeId string
			)
			if err = rows.Scan(&img.Id, &img.URL, &homeId); err != nil {
				return
			}

			home, ok := homesIds[homeId]
			if !ok {
				continue
			}

			home.Images = append(home.Images, img)
		}
	}

	for _, home := range homesIds {
		homes = append(homes, home)
	}

	return
}

func (h Home) getImage(batch *pgx.Batch, homeId string) {
	sql := `
	SELECT
	    id,
		link,
		homeid
	FROM
		home_images
	WHERE 
	    homeid=@homeid
	`

	arg := pgx.NamedArgs{
		"homeid": homeId,
	}

	batch.Queue(sql, arg)
}
