-- Table: public.orders

-- DROP TABLE public.orders;
CREATE TABLE public.orders
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    customer_id integer NOT NULL,
    vendor_id integer NOT NULL,
    location_id integer NOT NULL,
    status integer NOT NULL,
    total NUMERIC(6, 4) NOT NULL,
    tax NUMERIC(6, 4) NOT NULL,
    instructions text COLLATE pg_catalog."default",
    CONSTRAINT orders_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.orders
    OWNER to postgres;

-- Index: idx_orders_deleted_at

-- DROP INDEX public.idx_orders_deleted_at;

CREATE INDEX idx_orders_deleted_at
    ON public.orders USING btree
        (deleted_at)
    TABLESPACE pg_default;