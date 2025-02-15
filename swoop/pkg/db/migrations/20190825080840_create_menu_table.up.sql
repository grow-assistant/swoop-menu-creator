-- Table: public.locations

-- DROP TABLE public.locations;
CREATE TABLE public.menus
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    location_id integer NOT NULL,
    disabled boolean NOT NULL,
    CONSTRAINT menus_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.menus
    OWNER to postgres;

-- Index: idx_menus_deleted_at

-- DROP INDEX public.idx_menus_deleted_at;

CREATE INDEX idx_menus_deleted_at
    ON public.menus USING btree
        (deleted_at)
    TABLESPACE pg_default;