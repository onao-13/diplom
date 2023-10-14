-- HOME
CREATE TABLE cities(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE homes(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price INT NOT NULL,
    homeFeatureId BIGINT NOT NULL REFERENCES homes_features(id),
    cityId BIGINT NOT NULL REFERENCES cities(id)   
);

CREATE TABLE homes_features(
    homeId BIGINT NOT NULL REFERENCES homes(id),
    transportDataId BIGINT NOT NULL REFERENCES home_transport_data(id),
    popularLocationsId BIGINT NOT NULL REFERENCES home_popular_locations_data(id)
);

CREATE TABLE home_transport_data(
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL
);

CREATE TABLE home_popular_locations_data(
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL
);

-- ARTICLES
CREATE TABLE articles_categories(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE articles(
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    content VARCHAR(5000) NOT NULL,
    categotyId BIGINT NOT NULL REFERENCES articles_categories(id),
);
    -- imageId BIGINT NOT NULL REFERENCES images(id) 

-- CREATE TABLE images(
--     id SERIAL PRIMARY KEY
-- );
