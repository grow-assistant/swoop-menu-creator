-- Table: public.locations

ALTER TABLE public.locations
    ADD COLUMN timezone text;

UPDATE public.locations
    SET timezone = 'America/Denver'
    WHERE timezone IS NULL;

ALTER TABLE public.locations
    ALTER COLUMN timezone SET NOT NULL;





