-- +goose Up
-- +goose StatementBegin
ALTER TABLE attendances MODIFY COLUMN check_in_time DATETIME NULL;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE attendances MODIFY COLUMN check_out_time DATETIME NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE attendances MODIFY COLUMN check_in_time TIME NULL;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE attendances MODIFY COLUMN check_out_time TIME NULL;
-- +goose StatementEnd 