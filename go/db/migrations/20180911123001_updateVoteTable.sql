
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE votes ADD count int(32) DEFAULT 0 NOT NULL;
ALTER TABLE candidates ADD votes int(32) DEFAULT 0 NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

