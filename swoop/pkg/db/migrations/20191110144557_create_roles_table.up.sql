-- Table: public.roles

CREATE TABLE public.roles
(
    user_id integer NOT NULL,
    location_id integer NOT NULL,
    role_id integer NOT NULL
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.roles
    OWNER to postgres;

