CREATE TABLE IF NOT EXISTS Person(
    id VARCHAR(26) NOT NULL PRIMARY KEY,
    nickname VARCHAR(32) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    birthday date NOT NULL,
    stack VARCHAR(32)[]
);

CREATE INDEX IF NOT EXISTS person_fts_idx ON Person USING GIN (
    to_tsvector(nickname || ' ' || name || array_to_string(stack, ' ')));
