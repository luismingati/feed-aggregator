-- +goose Up

CREATE TABLE users (
  id UUID PRIMARY KEY,
  created_at timestamp not null,
  updated_at timestamp not null,
  name TEXT NOT NULL
);

-- +goose Down
drop table users;