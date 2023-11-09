-- +migrate Up
ALTER TABLE
    tono_command_permissions RENAME TO command_permissions;

-- +migrate Down
ALTER TABLE
    command_permissions RENAME TO tono_command_permissions;