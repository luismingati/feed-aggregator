-- +goose Up

CREATE TABLE Feeds(
  id UUID PRIMARY KEY,
  created_at timestamp not null,
  updated_at timestamp not null,
  name TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL,
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE Feeds;