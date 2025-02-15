-- Table: public.device_sessions

-- DROP TABLE public.device_sessions;

CREATE TABLE public.device_sessions
(
    id serial,
    device_token text COLLATE pg_catalog."default",
    device_type integer NOT NULL,
    user_id integer NOT NULL,
    expires_at timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT device_sessions_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.device_sessions
    OWNER to postgres;

-- Index: idx_device_sessions_deleted_at

-- DROP INDEX public.idx_device_sessions_deleted_at;

CREATE INDEX idx_device_sessions_deleted_at
    ON public.device_sessions USING btree
        (deleted_at)
    TABLESPACE pg_default;