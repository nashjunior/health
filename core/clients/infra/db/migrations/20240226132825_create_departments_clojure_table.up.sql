CREATE TABLE public.departments_closure(
    uuid uuid primary key,
    id serial4 not null unique,
    id_department int not null,
    id_manager int,
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp,

    CONSTRAINT "FK_DEPARTMENTS_CLOSURE_DEPARTMENT"
        FOREIGN KEY  (id_department)
        references public.departments(id)
        ON UPDATE CASCADE ON DELETE RESTRICT,

    CONSTRAINT "FK_DEPARTMENTS_CLOSURE_MANAGER"
        FOREIGN KEY  (id_manager)
        references public.departments(id)
        ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE INDEX idx_department_managers on public.departments_closure using hash(id_department);

