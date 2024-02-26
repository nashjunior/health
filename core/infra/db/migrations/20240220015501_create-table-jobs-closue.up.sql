CREATE TABLE public.jobs_closure(
    uuid uuid primary key,
    id serial4 not null unique,
    id_job int not null,
    id_manager int,
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp,

    CONSTRAINT "FK_JOBS_CLOSURE_JOB"
        FOREIGN KEY  (id_job)
        references public.jobs(id)
        ON UPDATE CASCADE ON DELETE RESTRICT,

    CONSTRAINT "FK_JOBS_CLOSURE_MANAGER"
        FOREIGN KEY  (id_manager)
        references public.jobs(id)
        ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE INDEX idx_job_transaction on public.jobs_closure using hash(id_job);

