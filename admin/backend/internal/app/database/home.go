package database

import (
	"admin/internal/app/payload"
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
	return Home{
		ctx:  ctx,
		pool: pool,
	}
}

func (h Home) Create(cityId string, home payload.Home) error {
	sql := `
	INSERT INTO
		city_homes(name, street, cityid, price, description, transports, popularLocations, layout, greenzone, infrastructure, events, schools)
	VALUES(@name, @street, @cityid, @price, @description, @transports, @popularLocations, @layout, @greenzone, @infrastructure, @events, @schools)
	RETURNING id
	`

	args := pgx.NamedArgs{
		"name":             home.Name,
		"street":           home.Street,
		"cityid":           cityId,
		"price":            home.Price,
		"description":      home.Description,
		"transports":       home.Transports,
		"popularLocations": home.PopularLocations,
		"layout":           home.Layout,
		"greenzone":        home.GreenZone,
		"infrastructure":   home.Infrastructure,
		"events":           home.Events,
		"schools":          home.Schools,
	}

	var (
		batch = &pgx.Batch{}
		id    string
	)

	if err := h.pool.QueryRow(h.ctx, sql, args).Scan(&id); err != nil {
		return err
	}

	if len(home.Images) != 0 {
		for _, image := range home.Images {
			h.batchAddImage(batch, id, image)
		}
	}

	br := h.pool.SendBatch(h.ctx, batch)
	defer br.Close()

	return nil
}

func (h Home) GetById(cityId, homeId string) (home *payload.Home, err error) {
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
	    cityid=@cityid
	AND
	    city_homes.id=@id
	`

	args := pgx.NamedArgs{
		"cityid": cityId,
		"id":     homeId,
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

func (h Home) GetByCityId(cityId string) (homes []*payload.Home, err error) {
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

	rows, err := h.pool.Query(h.ctx, sql, arg)
	defer rows.Close()
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

func (h Home) Update(homeId string, home payload.Home) error {
	sql := `
	UPDATE city_homes SET	
		name=@name, 
		street=@street,
		price=@price,
		description=@description,
		transports=@transports,
		popularLocations=@popularLocations,
		layout=@layout,
		greenzone=@greenzone,
		infrastructure=@infrastructure,
		events=@events,
		schools=@schools
	WHERE 
	    id=@id
	`

	var batch = &pgx.Batch{}

	args := pgx.NamedArgs{
		"name":             home.Name,
		"street":           home.Street,
		"price":            home.Price,
		"description":      home.Description,
		"transports":       home.Transports,
		"popularLocations": home.PopularLocations,
		"layout":           home.Layout,
		"greenzone":        home.GreenZone,
		"infrastructure":   home.Infrastructure,
		"events":           home.Events,
		"schools":          home.Schools,
		"id":               homeId,
	}

	if _, err := h.pool.Exec(h.ctx, sql, args); err != nil {
		return err
	}

	for _, img := range home.Images {
		h.batchAddImage(batch, homeId, img)
	}

	br := h.pool.SendBatch(h.ctx, batch)
	defer br.Close()

	return nil
}

func (h Home) Delete(cityId, homeId string) error {
	sql := `
	DELETE FROM 
	   city_homes
	WHERE
	    id=@id
	AND
	    cityid=@cityid
	`

	args := pgx.NamedArgs{
		"id":     homeId,
		"cityid": cityId,
	}

	if _, err := h.pool.Exec(h.ctx, sql, args); err != nil {
		return err
	}

	return nil
}

func (h Home) batchAddImage(batch *pgx.Batch, homeId string, img payload.HomeImage) {
	sql := `
	select insertOrUpdateHomeImageURL(@link, @imgid, @homeid);
	`

	args := pgx.NamedArgs{
		"imgid":  img.Id,
		"homeid": homeId,
		"link":   img.URL,
	}

	batch.Queue(sql, args)
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
