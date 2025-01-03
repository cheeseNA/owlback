-- create table
--     public.tasks (
--                      id uuid not null default gen_random_uuid (),
--                      created_at timestamp with time zone null,
--                      created_by uuid not null,
--                      updated_at timestamp with time zone null,
--                      site_url text not null,
--                      condition_query text not null,
--                      duration_day bigint not null,
--                      is_public boolean not null,
--                      deleted_at timestamp with time zone null,
--                      is_paused boolean not null,
--                      constraint tasks_pkey primary key (id)
-- ) tablespace pg_default;