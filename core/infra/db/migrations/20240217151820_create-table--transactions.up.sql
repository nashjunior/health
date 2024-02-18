CREATE TABLE accounting.transactions(
    uuid uuid primary key,
    id serial4 not null unique,
    date timestamp not null,
    value numeric(10,2) not null,
    id_type_transaction int not null,
    created_at timestamp not null default now(),
    updated_at timestamp,
    deleted_at timestamp,

    CONSTRAINT "FK_TRANSACTIONS_TYPE_TRANSACTIONS"
        FOREIGN KEY  (id_type_transaction)
        references accounting.types_transactions(id)
        ON UPDATE CASCADE ON DELETE RESTRICT

);

CREATE INDEX idx_transaction_type_transaction on accounting.transactions using hash(id_type_transaction);
