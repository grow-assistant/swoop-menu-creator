
-- The purpose of this migration is to bump the id's
-- of the role types so that we do not use 0 as the initial
-- id because it causes some weird behavior
DELETE FROM public.role_types;

INSERT INTO public.role_types (id, name, created_at, updated_at) VALUES (1, 'Admin', Now(), Now());
INSERT INTO public.role_types (id, name, created_at, updated_at) VALUES (2, 'Runner', Now(), Now());
INSERT INTO public.role_types (id, name, created_at, updated_at) VALUES (3, 'Station', Now(), Now());