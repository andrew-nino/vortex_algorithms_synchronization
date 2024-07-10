CREATE TABLE IF NOT EXISTS managers (
    id            serial       primary key,
    name          varchar(255) not null,
    managername   varchar(255) not null unique,
    password_hash varchar(255) not null,
    role          varchar(7)   not null default 'manager',
    created_at    timestamp    not null default now()

    CONSTRAINT role_manager CHECK (role IN ('admin', 'manager'))
);

CREATE TABLE IF NOT EXISTS clients
(
    id              serial       not null,
    client_id       integer      not null unique,
    client_name     varchar(255) not null,
    version         integer      not null,
    image           varchar(255) not null,
    cpu             varchar(255) not null,
    memory          varchar(255) not null,
    priority        numeric(5,2) not null,
    needRestart     boolean      not null default false,
    spawned_at      varchar(255) not null,
    created_at      timestamp    not null default now(),
    update_at       timestamp    not null default now()
);

CREATE UNIQUE INDEX clients_name_on_id_idx ON clients (client_id, client_name);

CREATE TABLE IF NOT EXISTS algorithm_status(
    id          serial  NOT NULL,
    client_id   integer NOT NULL references clients(client_id),
    VWAP        boolean NOT NULL default false,
    TWAP        boolean NOT NULL default false,
    HFT         boolean NOT NULL default false
);