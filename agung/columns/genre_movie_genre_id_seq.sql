create sequence genre_movie_genre_id_seq;

alter sequence genre_movie_genre_id_seq owner to postgres;

alter sequence genre_movie_genre_id_seq owned by genre_movie.genre_id;

