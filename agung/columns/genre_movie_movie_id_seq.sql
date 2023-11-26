create sequence genre_movie_movie_id_seq;

alter sequence genre_movie_movie_id_seq owner to postgres;

alter sequence genre_movie_movie_id_seq owned by genre_movie.movie_id;

