-- Table: public.locations

ALTER TABLE public.locations
    ADD COLUMN active boolean NOT NULL default true;

ALTER TABLE public.locations
    ADD COLUMN disabled boolean NOT NULL default false;








