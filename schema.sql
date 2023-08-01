CREATE TABLE if not exists todos (
  id text PRIMARY KEY,
  title text NOT NULL,
  is_done boolean NOT NULL
);