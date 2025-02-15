
-- Make sure we delete previous role types
-- to maintain id sequencing when we insert
DELETE FROM public.role_types;

INSERT INTO public.role_types (id, name, created_at, updated_at) VALUES (0, 'Admin', Now(), Now());
INSERT INTO public.role_types (id, name, created_at, updated_at) VALUES (1, 'Runner', Now(), Now());
INSERT INTO public.role_types (id, name, created_at, updated_at) VALUES (2, 'Station', Now(), Now());