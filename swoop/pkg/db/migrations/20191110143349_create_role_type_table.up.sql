-- Table: public.role_types

-- DROP TABLE public.role_types;

CREATE TABLE public.role_types
(
    id serial,
    name text COLLATE pg_catalog."default",
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT role_types_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.role_types
    OWNER to postgres;

-- Index: idx_role_types_deleted_at

-- DROP INDEX public.idx_role_types_deleted_at;

CREATE INDEX idx_role_types_deleted_at
    ON public.role_types USING btree
        (deleted_at)
    TABLESPACE pg_default;