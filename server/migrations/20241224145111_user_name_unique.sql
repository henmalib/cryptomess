-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD UNIQUE(username)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP CONSTRAINT username;
-- +goose StatementEnd
