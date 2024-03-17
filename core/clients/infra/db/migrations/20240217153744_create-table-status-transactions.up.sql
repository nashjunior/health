CREATE TABLE accounting.status_transactions(
     uuid uuid primary key,
    id serial4 not null unique,
    status smallint not null,
    id_transaction int not null,
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp,

    CONSTRAINT "FK_STATUSTRANSACTIONS_TRANSACTION"
        FOREIGN KEY  (id_transaction)
        references accounting.transactions(id)
        ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE INDEX idx_status_transactions_transaction on accounting.status_transactions using hash(id_transaction);
