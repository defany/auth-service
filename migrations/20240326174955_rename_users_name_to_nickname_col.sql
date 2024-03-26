-- +goose Up
-- +goose StatementBegin
alter table users rename column name to nickname;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users rename column nickname to name;
-- +goose StatementEnd
