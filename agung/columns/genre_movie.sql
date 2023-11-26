create table genre_movie
(
    genre_id bigint not null
        constraint genre_movie_genres_id_fk
            references genres,
    movie_id bigint not null
        constraint genre_movie_movies_id_fk
            references movies
);

alter table genre_movie
    owner to postgres;

