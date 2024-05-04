-- HOME
CREATE TABLE cities
(
    id   BIGSERIAL  NOT NULL,
    name VARCHAR(100) NOT NULL UNIQUE,
    PRIMARY KEY (id)
);

COMMENT ON TABLE cities IS 'Города';

CREATE TABLE city_homes
(
    id     BIGSERIAL  NOT NULL,
    name   VARCHAR(200) NOT NULL,
    street VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    cityId BIGINT  NOT NULL,
    price  INT     NOT NULL DEFAULT 0,
    layout VARCHAR(500),
    greenzone VARCHAR(500),
    infrastructure VARCHAR(500),
    events VARCHAR(500),
    schools VARCHAR(500),
    transports VARCHAR(500),
    popularLocations VARCHAR(500),
    PRIMARY KEY (id)
);

COMMENT ON TABLE city_homes IS 'Список домов в городе';

CREATE TABLE home_images
(
    id BIGSERIAL PRIMARY KEY,
    homeId BIGINT  NOT NULL,
    link   VARCHAR(100) NOT NULL
);

COMMENT ON TABLE home_images IS 'Картинки к домам';

create table manager_call
(
    id     bigserial
        constraint manager_call_pk
            primary key,
    name   varchar(100) not null,
    number varchar(20)  not null,
    homeId bigint       not null
        constraint manager_call_city_homes_id_fk
            references city_homes
);

create table users(
    id bigserial primary key,
    username varchar(100) not null,
    password varchar(100) not null
);

alter table city_homes
    add constraint fk_citites_to_city_homes
        foreign key (cityid) references cities
            on update cascade on delete cascade;

alter table home_images
    add constraint fk_city_homes_to_home_images
        foreign key (homeid) references city_homes
            on delete cascade;

create function insertOrUpdateHomeImageURL(newLink VARCHAR, imgId BIGINT, homeIdIns BIGINT)
returns INT
language plpgsql
as
$$
declare
    img_id bigint;
begin
    select id into img_id
    from home_images
    where id=imgId;

    if not exists(select id from home_images where id=imgId) then
        INSERT INTO
            home_images(homeId, link)
        VALUES(homeIdIns, newLink);
    else
        UPDATE
            home_images
        SET
            link=newLink
        WHERE
            id=imgId;
    end if;
    return 0;
end;
$$
