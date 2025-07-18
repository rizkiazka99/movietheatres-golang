-- +migrate Up
-- +migrate StatementBegin

create table theatres (
    id BIGINT PRIMARY KEY NOT NULL,
    nama varchar(128) NOT NULL,
    lokasi varchar(500) NOT NULL,
    rating DOUBLE PRECISION NOT NULL
)

-- +migrate StatementEnd