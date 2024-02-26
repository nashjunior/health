CREATE TABLE public.jobs(
     uuid uuid primary key,
    id serial4 not null unique,
    name varchar(36) not null,
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE INDEX idx_job_is_active on public.jobs(deleted_at) WHERE deleted_at IS NOT NULL;
