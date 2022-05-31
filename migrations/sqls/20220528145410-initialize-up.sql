/* Replace with your SQL commands */

-- +migrate Up
create table "nft_collections" (
"address" text primary key,
"name" text,
"symbol" text,
"chain_id" integer,
"erc_format" text
);