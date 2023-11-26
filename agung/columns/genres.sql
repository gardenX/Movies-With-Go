create table genres
(
    id         bigserial
        constraint genres_pk
            primary key,
    name       varchar(120) not null,
    created_at timestamp    not null,
    updated_at timestamp    not null,
    deleted_at timestamp
);

alter table genres
    owner to postgres;

