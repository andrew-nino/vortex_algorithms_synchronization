CREATE TABLE IF NOT EXISTS algorithm_status(
    id          serial  NOT NULL,
    client_id   integer NOT NULL unique,
    VWAP        boolean NOT NULL default false,
    TWAP        boolean NOT NULL default false,
    HFT         boolean NOT NULL default false
);