ALTER TABLE public.markers
    ALTER COLUMN disabled SET DEFAULT false;

ALTER TABLE public.markers
    ALTER COLUMN disabled SET NOT NULL;