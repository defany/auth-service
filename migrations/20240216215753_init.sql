-- +goose Up
-- +goose StatementBegin
create type user_role as Enum('unknown', 'user', 'admin');

create table if not exists users(
    id bigserial primary key constraint positive_user_id check ( id > 0 ),
    name text not null,
    email text,
    password text not null,
    password_confirm text not null,
    role user_role not null,
    created_at timestamp not null default clock_timestamp(),
    updated_at timestamp not null default clock_timestamp()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
