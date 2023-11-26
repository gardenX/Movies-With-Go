create table people
(
    id         bigserial
        constraint people_pk
            primary key,
    name       varchar   not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp
);

alter table people
    owner to postgres;

