-- +migrate Up
ALTER TABLE auto_condition_values RENAME COLUMN type TO type_id;
ALTER TABLE auto_condition_types RENAME TO auto_types;
alter table auto_actions drop column embed_id;
-- +migrate Down
ALTER TABLE auto_condition_values RENAME COLUMN type_id TO type;
ALTER TABLE auto_types RENAME TO auto_condition_types;
alter table auto_actions add column embed_id text;