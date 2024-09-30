-- +goose Up
-- +goose StatementBegin
CREATE TABLE auth (

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth;
-- +goose StatementEnd