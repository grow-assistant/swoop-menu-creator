-- Table: public.options

-- DROP TABLE public.options;
CREATE TABLE public.options
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    min integer NOT NULL,
    max integer NOT NULL,
    item_id integer NOT NULL,
    disabled boolean NOT NULL,
    CONSTRAINT options_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.options
    OWNER to postgres;

-- Index: idx_options_deleted_at

-- DROP INDEX public.idx_options_deleted_at;

CREATE INDEX idx_options_deleted_at
    ON public.options USING btree
        (deleted_at)
    TABLESPACE pg_default;