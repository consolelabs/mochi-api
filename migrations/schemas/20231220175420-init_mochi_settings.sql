
-- +migrate Up
CREATE SEQUENCE IF NOT EXISTS user_payment_settings_id_seq;
CREATE TABLE IF NOT EXISTS user_payment_settings (
	id int4 not null default nextval('user_payment_settings_id_seq'::regclass),
	profile_id text not null unique,
	default_money_source jsonb,
	default_receiver_platform text not null default '',
	default_token_id text,
	prioritized_token_ids text[] not null default array[]::varchar[],
	default_message_enable boolean not null default false,
	default_message_settings jsonb,
	tx_limit_enable boolean not null default false,
	tx_limit_settings jsonb
);

CREATE SEQUENCE IF NOT EXISTS user_privacy_settings_id_seq;
CREATE TABLE IF NOT EXISTS user_privacy_settings (
	id int4 not null default nextval('user_privacy_settings_id_seq'::regclass),
	profile_id text not null unique,
	tx jsonb not null,
	social_accounts jsonb not null,
	wallets jsonb not null
);

CREATE TABLE IF NOT EXISTS notification_flags (
	"key" text primary key,
	"group" text not null,
	description text not null
);

CREATE SEQUENCE IF NOT EXISTS user_notification_settings_id_seq;
CREATE TABLE IF NOT EXISTS user_notification_settings (
	id int4 not null default nextval('user_notification_settings_id_seq'::regclass),
	profile_id text not null unique,
	enable boolean not null default true,
	platforms text[] not null default array[]::varchar[],
	flags jsonb not null
);

-- +migrate Down
DROP TABLE IF EXISTS user_notification_settings;
DROP TABLE IF EXISTS notification_flags;
DROP TABLE IF EXISTS user_privacy_settings;
DROP TABLE IF EXISTS user_payment_settings;

DROP SEQUENCE IF EXISTS user_notification_settings_id_seq;
DROP SEQUENCE IF EXISTS user_privacy_settings_id_seq;
DROP SEQUENCE IF EXISTS user_payment_settings_id_seq;