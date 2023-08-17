CREATE TABLE IF NOT EXISTS Person(
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    nickname VARCHAR(32) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    birthday VARCHAR(10) NOT NULL,
    stack VARCHAR(32)[]
);

CREATE OR REPLACE FUNCTION f_stack_array(varchar[])
  RETURNS text LANGUAGE sql IMMUTABLE AS $$SELECT array_to_string($1, ' ')$$;

ALTER TABLE Person
    ADD COLUMN IF NOT EXISTS docsearch tsvector
        GENERATED ALWAYS AS (
            to_tsvector(
                'portuguese',
                nickname || ' ' || name || ' ' || COALESCE(f_stack_array(stack), ''))) STORED;

CREATE INDEX IF NOT EXISTS docsearch_idx ON Person USING GIN(docsearch);
