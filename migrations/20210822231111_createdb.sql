-- +goose Up
-- +goose StatementBegin
CREATE TABLE vacancies (
    id      SERIAL PRIMARY KEY,
    link    TEXT,
    status  INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE vacancies;
-- +goose StatementEnd
