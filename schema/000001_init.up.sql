CREATE TABLE IF NOT EXISTS managers (
    id            serial       primary key,
    name          varchar(255) not null,
    managername   varchar(255) not null unique,
    password_hash varchar(255) not null,
    role          varchar(7)   not null default 'manager',
    created_at    timestamp    not null default now()

    CONSTRAINT role_manager CHECK (role IN ('admin', 'manager'))
);

CREATE TABLE IF NOT EXISTS algorithm_status(
    id          serial  NOT NULL,
    client_id   integer NOT NULL unique,
    VWAP        boolean NOT NULL default false,
    TWAP        boolean NOT NULL default false,
    HFT         boolean NOT NULL default false
);