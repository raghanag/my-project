CREATE SCHEMA weather_ts;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS timescaledb;
CREATE TABLE IF NOT EXISTS weather_ts.minutely_weather
(
    id uuid DEFAULT uuid_generate_v4 (),
    dt TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    data_type  text  NOT NULL,
    value  decimal  NOT NULL,
    PRIMARY KEY (id, dt)
);
SELECT create_hypertable('weather_ts.minutely_weather','dt',create_default_indexes => FALSE);