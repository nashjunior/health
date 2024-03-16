CREATE TABLE accounting.types_transactions(
    uuid uuid primary key,
    id serial4 not null unique,
    name varchar(36) not null,
    account_type smallint,
    operation_type smallint,
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE INDEX idx_account_type on accounting.types_transactions using hash(account_type);
