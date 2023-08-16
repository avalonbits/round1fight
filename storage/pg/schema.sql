CREATE TABLE IF NOT EXISTS Person(
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    nickname VARCHAR(32) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    birthday date NOT NULL,
    stack VARCHAR(32)[]
);

CREATE OR REPLACE FUNCTION f_stack_array(varchar[])
  RETURNS text LANGUAGE sql IMMUTABLE AS $$SELECT array_to_string($1, ' ')$$;

CREATE INDEX IF NOT EXISTS person_fts_idx ON Person USING GIN (
    to_tsvector('portuguese', nickname || ' ' || name || COALESCE(f_stack_array(stack), '')));
