-- +goose Up
-- +goose StatementBegin
create table if not exists roles_endpoints_permissions(
    role text not null,
    endpoint text not null,

    primary key (role, endpoint)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists roles_endpoints_permissions;
-- +goose StatementEnd
