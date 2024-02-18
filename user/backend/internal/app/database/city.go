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

	city.Homes = make([]*payload.Home, 0)

	rows, err := c.pool.Query(c.ctx, sql, arg)
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

		city.Homes = append(city.Homes, &home)
	}
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
