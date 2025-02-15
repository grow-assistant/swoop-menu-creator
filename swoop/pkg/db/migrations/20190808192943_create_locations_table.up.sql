-- Table: public.locations

-- DROP TABLE public.locations;
CREATE TABLE public.locations
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    CONSTRAINT locations_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.locations
    OWNER to postgres;

-- Index: idx_locations_deleted_at

-- DROP INDEX public.idx_locations_deleted_at;

CREATE INDEX idx_locations_deleted_at
    ON public.locations USING btree
    (deleted_at)
    TABLESPACE pg_default;