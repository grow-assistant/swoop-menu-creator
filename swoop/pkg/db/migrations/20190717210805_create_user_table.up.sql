-- Table: public.users

-- DROP TABLE public.users;

CREATE TABLE public.users
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;

-- Index: idx_users_deleted_at

-- DROP INDEX public.idx_users_deleted_at;

CREATE INDEX idx_users_deleted_at
    ON public.users USING btree
    (deleted_at)
    TABLESPACE pg_default;