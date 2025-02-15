-- Table: public.categories

-- DROP TABLE public.categories;
CREATE TABLE public.categories
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    description text COLLATE pg_catalog."default",
    menu_id integer NOT NULL,
    disabled boolean NOT NULL,
    CONSTRAINT categories_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.categories
    OWNER to postgres;

-- Index: idx_categories_deleted_at

-- DROP INDEX public.idx_categories_deleted_at;

CREATE INDEX idx_categories_deleted_at
    ON public.categories USING btree
        (deleted_at)
    TABLESPACE pg_default;