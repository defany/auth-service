-- +goose Up
-- +goose StatementBegin
create table if not exists logs(
    action text not null,
    user_id bigserial references users(id),
    timestamp timestamp not null default clock_timestamp()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists logs;
-- +goose StatementEnd
