-- +goose Up
create table messages
(
    id         serial primary key,
    text       text,
    username varchar,
    chat_id int not null ,
    created_at timestamp not null default now()
);

alter table messages
    add constraint messages_chat_id_fkey
        foreign key (chat_id) references chats
            on delete cascade;

-- +goose Down
alter table messages drop constraint messages_chat_id_fkey;
drop table messages;

