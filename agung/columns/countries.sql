create table countries
(
    id         bigserial
        constraint countries_pk
            primary key,
    name       varchar(120) not null,
    created_at timestamp    not null,
    updated_at timestamp    not null,
    deleted_at timestamp
);

alter table countries
    owner to postgres;

