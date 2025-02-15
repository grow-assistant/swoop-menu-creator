-- Table: public.identities

-- DROP TABLE public.identities;
CREATE TABLE public.identities
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    provider text COLLATE pg_catalog."default",
    uid text COLLATE pg_catalog."default",
    user_id integer NOT NULL,
    first_name text COLLATE pg_catalog."default",
    last_name text COLLATE pg_catalog."default",
    email text COLLATE pg_catalog."default",
    picture text COLLATE pg_catalog."default",
    CONSTRAINT identities_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.identities
    OWNER to postgres;

-- Index: idx_identities_deleted_at

-- DROP INDEX public.idx_identities_deleted_at;

CREATE INDEX idx_identities_deleted_at
    ON public.identities USING btree
    (deleted_at)
    TABLESPACE pg_default;