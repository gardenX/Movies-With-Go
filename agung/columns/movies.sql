create table movies
(
    id          bigserial
        constraint movies_pk
            primary key,
    title       varchar   not null,
    released_at date      not null,
    created_at  timestamp not null,
    updated_at  timestamp not null,
    deleted_at  timestamp,
    director_id bigint
        constraint movies_director_id_fk
            references people,
    country_id  bigint
        constraint movies_countries_id_fk
            references countries
);

alter table movies
    owner to postgres;

