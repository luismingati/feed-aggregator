-- +goose Up

CREATE TABLE posts (
  id UUID PRIMARY KEY,
  created_at timestamp not null,
  updated_at timestamp not null,
  title TEXT NOT NULL,
  description TEXT,
  published_at TIMESTAMP NOT NULL,
  url text unique not null,
  feed_id UUID NOT NULL REFERENCES Feeds(id) ON DELETE CASCADE
);

-- +goose Down
drop table posts;