CREATE TABLE public.departments(
     uuid uuid primary key,
    id serial4 not null unique,
    name varchar(36) not null,
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE INDEX idx_department_is_active on public.departments(deleted_at) WHERE deleted_at IS NOT NULL;
