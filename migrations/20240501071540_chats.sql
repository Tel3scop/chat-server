-- +goose Up
create table chats
(
    id         serial primary key,
    name       varchar,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

-- +goose Down
drop table chats;

