create table if not exists movies
(
    id         bigserial
        primary key,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    title      text,
    director   text,
    year       text,
    plot       text
);

create index if not exists idx_movies_deleted_at
    on movies (deleted_at);

