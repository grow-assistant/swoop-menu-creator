-- Table: public.items

-- DROP TABLE public.items;
CREATE TABLE public.items
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    price NUMERIC(6, 4) NOT NULL,
    category_id integer NOT NULL,
    disabled boolean NOT NULL,
    CONSTRAINT items_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.items
    OWNER to postgres;

-- Index: idx_items_deleted_at

-- DROP INDEX public.idx_items_deleted_at;

CREATE INDEX idx_items_deleted_at
    ON public.items USING btree
        (deleted_at)
    TABLESPACE pg_default;