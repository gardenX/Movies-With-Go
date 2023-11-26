create sequence movies_director_id_seq;

alter sequence movies_director_id_seq owner to postgres;

alter sequence movies_director_id_seq owned by movies.director_id;

