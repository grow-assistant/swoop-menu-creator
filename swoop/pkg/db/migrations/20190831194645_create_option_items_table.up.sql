-- Table: public.option_items

-- DROP TABLE public.option_items;
CREATE TABLE public.option_items
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    price NUMERIC(6, 4) NOT NULL,
    option_id integer NOT NULL,
    disabled boolean NOT NULL,
    CONSTRAINT option_items_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.option_items
    OWNER to postgres;

-- Index: idx_option_items_deleted_at

-- DROP INDEX public.idx_option_items_deleted_at;

CREATE INDEX idx_option_items_deleted_at
    ON public.option_items USING btree
        (deleted_at)
    TABLESPACE pg_default;