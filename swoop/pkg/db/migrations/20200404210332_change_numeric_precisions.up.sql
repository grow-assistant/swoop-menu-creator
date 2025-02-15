ALTER table public.orders
    ALTER COLUMN total TYPE numeric(6,2);

ALTER table public.option_items
    ALTER COLUMN price TYPE numeric(6,2);