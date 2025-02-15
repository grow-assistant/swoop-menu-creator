-- Table: public.discounts

-- DROP TABLE public.discounts;
CREATE TABLE public.discounts
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    order_id integer NOT NULL,
    user_id integer NOT NULL,
    amount NUMERIC(6, 4) NOT NULL,
    reason text COLLATE pg_catalog."default",
    CONSTRAINT discounts_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.discounts
    OWNER to postgres;

-- Index: idx_discounts_deleted_at

-- DROP INDEX public.idx_discounts_deleted_at;

CREATE INDEX idx_discounts_deleted_at
    ON public.discounts USING btree
        (deleted_at)
    TABLESPACE pg_default;