-- Table: public.user_locations

CREATE TABLE public.user_locations
(
    user_id integer NOT NULL,
    latitude float NOT NULL,
    longitude float NOT NULL
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.user_locations
    OWNER to postgres;

