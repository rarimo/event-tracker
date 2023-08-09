-- +migrate Up

create table processed_txs
(
    id    bigserial primary key,
    block bigint      not null,
    hash  text unique not null
);

-- +migrate Down

drop table processed_txs;
