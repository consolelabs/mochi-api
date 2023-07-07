
-- +migrate Up
alter table treasurers rename to vault_treasurers;
alter table treasurer_requests rename to vault_requests;
alter table treasurer_submissions rename to vault_submissions;
-- +migrate Down
alter table vault_treasurers rename to treasurers;
alter table vault_requests rename to treasurer_requests;
alter table vault_submissions rename to treasurer_submissions;
