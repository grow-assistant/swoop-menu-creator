-- Table: public.markers

-- DROP TABLE public.markers
CREATE TABLE public.markers
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    disabled boolean,
    CONSTRAINT markers_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.markers
    OWNER to postgres;

-- Index: idx_markers_deleted_at

-- DROP INDEX public.idx_markers_deleted_at;

CREATE INDEX idx_markers_deleted_at
    ON public.markers USING btree
    (deleted_at)
    TABLESPACE pg_default;