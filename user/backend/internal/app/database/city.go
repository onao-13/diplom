package database

import (
	"backend/internal/app/payload"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type City struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func NewCity(ctx context.Context, pool *pgxpool.Pool) City {
	return City{ctx: ctx, pool: pool}
}

func (c City) GetAll() (cities []*payload.City, err error) {
	cities, err = c.getCities()
	if err != nil {
		return
	}

	var wg = &sync.WaitGroup{}

	for _, city := range cities {
		wg.Add(1)
		c.getHomesByCityId(wg, city.Id, city)
	}

	wg.Wait()

	return
}

func (c City) getCities() (cities []*payload.City, err error) {
	sql := `
	SELECT
		id,
		name
	FROM
	    cities
	`

	rows, err := c.pool.Query(c.ctx, sql)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var city payload.City
		if err := rows.Scan(&city.Id, &city.Name); err != nil {
			break
		}

		cities = append(cities, &city)
	}

	return
}

func (c City) getHomesByCityId(wg *sync.WaitGroup, cityId string, city *payload.City) {
	defer wg.Done()

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
	`

	arg := pgx.NamedArgs{
		"cityid": cityId,
	}

	city.Homes = make([]*payload.Home, 0)

	rows, err := c.pool.Query(c.ctx, sql, arg)
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
		c.getImage(batch, home.Id)
	}

	br := c.pool.SendBatch(c.ctx, batch)
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
		city.Homes = append(city.Homes, home)
	}
}

func (c City) getImage(batch *pgx.Batch, homeId string) {
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

//func (c City) GetAll() (cities []*payload.City, err error) {
//	sql := `
//	SELECT
//	    id,
//	    name
//	FROM
//		 cities
//	`
//
//	rowsCity, err := c.pool.Query(c.ctx, sql)
//	if err != nil {
//		return
//	}
//
//	var batch = &pgx.Batch{}
//
//	// получение списка городов
//	for rowsCity.Next() {
//		var city payload.City
//		if err := rowsCity.Scan(&city.Id, &city.Name); err != nil {
//			break
//		}
//
//		c.batchGetCityHomes(batch, city.Id)
//		cities
//	}
//
//	defer rowsCity.Close()
//
//	br := c.pool.SendBatch(c.ctx, batch)
//
//	defer br.Close()
//
//	// получение списка домов в городе
//	for i := 0; i < batch.Len(); i++ {
//		rowsHomes, err := br.Query()
//		if err != nil {
//			continue
//		}
//
//		for rowsHomes.Next() {
//			var home payload.Home
//			err = rowsHomes.Scan(
//				&home.Id, &home.Name, &home.Street,
//				&home.Price, &home.CityId)
//			if err != nil {
//				break
//			}
//
//			city, ok := cityMap[home.CityId]
//			if !ok {
//				continue
//			}
//
//			city.Homes = append(city.Homes, home)
//		}
//	}
//
//	cities = make([]payload.City, 0, len(cityMap))
//	for _, city := range cityMap {
//		cities = append(cities, *city)
//	}
//
//	return
//}
//
//func (c City) batchGetCityHomes(batch *pgx.Batch, cityId string) {
//	sql := `
//	SELECT
//		id,
//		name,
//		street,
//		price,
//		cityId
//	FROM
//		 city_homes
//	WHERE
//		cityId = @cityId
//	LIMIT 6
//	`
//
//	arg := pgx.NamedArgs{
//		"cityId": cityId,
//	}
//
//	batch.Queue(sql, arg)
//}
