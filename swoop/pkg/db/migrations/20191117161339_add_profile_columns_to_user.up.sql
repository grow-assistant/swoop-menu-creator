-- Table: public.users

ALTER TABLE public.users
    ADD COLUMN first_name text COLLATE pg_catalog."default";

ALTER TABLE public.users
    ADD COLUMN last_name text COLLATE pg_catalog."default";

ALTER TABLE public.users
    ADD COLUMN email text COLLATE pg_catalog."default";

