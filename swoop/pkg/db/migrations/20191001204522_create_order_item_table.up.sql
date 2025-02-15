-- Table: public.order_items

-- DROP TABLE public.order_items;
CREATE TABLE public.order_items
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    item_id integer NOT NULL,
    quantity integer NOT NULL,
    order_id integer NOT NULL,
    uid text COLLATE pg_catalog."default",
    CONSTRAINT order_items_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.order_items
    OWNER to postgres;

-- Index: idx_order_items_deleted_at

-- DROP INDEX public.idx_order_items_deleted_at;

CREATE INDEX idx_order_items_deleted_at
    ON public.order_items USING btree
        (deleted_at)
    TABLESPACE pg_default;