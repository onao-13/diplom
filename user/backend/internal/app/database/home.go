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

func (h *Home) GetById(id int64) (home *payload.Home, err error) {
	sql := `
	SELECT
		city_homes.id,
		city_homes.name,
		city_homes.street,
		city_homes.price,
		city_homes.cityid,
		COALESCE(city_homes.layout, ''),
		COALESCE(city_homes.greenzone, ''),
		COALESCE(city_homes.infrastructure, ''),
		COALESCE(city_homes.events, ''),
		COALESCE(city_homes.schools, ''),
		COALESCE(home_transports.id, 0),
		COALESCE(home_transports.name, ''),
		COALESCE(home_popular_locations.id, 0),
		COALESCE(home_popular_locations.name, ''),
		COALESCE(home_popular_locations.address, ''),
		COALESCE(home_images.link, '')
	FROM city_homes
		LEFT JOIN home_transports ON city_homes.id = home_transports.homeid
		LEFT JOIN home_popular_locations ON city_homes.id = home_popular_locations.homeid
		LEFT JOIN home_images on city_homes.id = home_images.homeid
	WHERE 
		city_homes.id=$1
	`

	rows, err := h.pool.Query(h.ctx, sql, id)
	if err != nil {
		return
	}

	home = &payload.Home{}
	for rows.Next() {
		var transport payload.HomeTransport
		var location payload.HomePopularLocation
		var img payload.HomeImage

		err = rows.Scan(
			&home.Id, &home.Name, &home.Street, &home.Price, &home.CityId,
			&home.Layout, &home.GreenZone, &home.Infrastructure, &home.Events,
			&home.Schools,
			&transport.Id, &transport.Name,
			&location.Id, &location.Name, &location.Address,
			&img.URL,
		)
		if err != nil {
			return
		}

		home.Images = append(home.Images, img)
		home.PopularLocations = append(home.PopularLocations, location)
		home.Transports = append(home.Transports, transport)
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
		COALESCE(city_homes.layout, ''),
		COALESCE(city_homes.greenzone, ''),
		COALESCE(city_homes.infrastructure, ''),
		COALESCE(city_homes.events, ''),
		COALESCE(city_homes.schools, ''),
		COALESCE(home_transports.id, 0),
		COALESCE(home_transports.name, ''),
		COALESCE(home_popular_locations.id, 0),
		COALESCE(home_popular_locations.name, ''),
		COALESCE(home_popular_locations.address, ''),
		COALESCE(home_images.link, '')
	FROM city_homes
		LEFT JOIN home_transports ON city_homes.id = home_transports.homeid
		LEFT JOIN home_popular_locations ON city_homes.id = home_popular_locations.homeid
		LEFT JOIN home_images on city_homes.id = home_images.homeid
	WHERE 
	    cityid=@cityid
	`

	arg := pgx.NamedArgs{
		"cityid": cityId,
	}

	rows, err := h.pool.Query(h.ctx, sql, arg)
	if err != nil {
		return
	}

	for rows.Next() {
		var home payload.Home
		var transport payload.HomeTransport
		var location payload.HomePopularLocation
		var img payload.HomeImage

		err = rows.Scan(
			&home.Id, &home.Name, &home.Street, &home.Price, &home.CityId,
			&home.Layout, &home.GreenZone, &home.Infrastructure, &home.Events,
			&home.Schools,
			&transport.Id, &transport.Name,
			&location.Id, &location.Name, &location.Address,
			&img.URL,
		)
		if err != nil {
			fmt.Println("Err scan: ", err)
			return
		}

		home.Transports = append(home.Transports, transport)
		home.PopularLocations = append(home.PopularLocations, location)
		home.Images = append(home.Images, img)

		homes = append(homes, &home)
	}

	return
}
