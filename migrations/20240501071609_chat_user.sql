-- +goose Up
create table chat_user
(
    chat_id         int not null,
    username varchar not null,
    created_at timestamp not null default now()
);

create unique index chat_user_uq
    on chat_user (chat_id, username);

alter table chat_user
    add constraint chat_user_chat_id_fkey
        foreign key (chat_id) references chats
            on delete cascade;



-- +goose Down
drop index chat_user_uq;
alter table chat_user drop constraint chat_user_chat_id_fkey;
drop table chat_user;

