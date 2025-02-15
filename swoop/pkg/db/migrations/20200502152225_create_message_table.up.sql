-- Table: public.messages

-- DROP TABLE public.messages;
CREATE TABLE public.messages
(
    id serial,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    content text COLLATE pg_catalog."default",
    sender_id integer NOT NULL,
    recipient_id integer NOT NULL,
    order_id interval NOT NULL ,
    CONSTRAINT messages_pkey PRIMARY KEY (id)
)
    WITH (
        OIDS = FALSE
    )
    TABLESPACE pg_default;

ALTER TABLE public.messages
    OWNER to postgres;

-- Index: idx_messages_deleted_at

-- DROP INDEX public.idx_messages_deleted_at;

CREATE INDEX idx_messages_deleted_at
    ON public.messages USING btree
        (deleted_at)
    TABLESPACE pg_default;