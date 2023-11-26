create sequence movies_country_id_seq;

alter sequence movies_country_id_seq owner to postgres;

alter sequence movies_country_id_seq owned by movies.country_id;

