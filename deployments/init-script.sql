-- HOME
CREATE TABLE cities(
    id SERIAL PRIMARY KEY,
    cityName VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE homes(
    id SERIAL PRIMARY KEY,
    homeName VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    homeFeatureId BIGINT NOT NULL REFERENCES homes_features(id),
    cityId BIGINT NOT NULL REFERENCES cities(id)   
);

CREATE TABLE homes_features(
    homeId BIGINT NOT NULL REFERENCES homes(id),
    transportDataId BIGINT NOT NULL REFERENCES home_transport_data(id),
    popularLocationsDataId BIGINT NOT NULL REFERENCES home_popular_locations_data(id)
);

CREATE TABLE home_transport_data(
    id SERIAL PRIMARY KEY,
    transportName VARCHAR(150) NOT NULL
);

CREATE TABLE home_popular_locations_data(
    id SERIAL PRIMARY KEY,
    popularLocationsName VARCHAR(150) NOT NULL
);

-- ARTICLES
CREATE TABLE artciles(
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(5000) NOT NULL
);
    -- imageId BIGINT NOT NULL REFERENCES images(id) 

-- CREATE TABLE images(
--     id SERIAL PRIMARY KEY
-- );
