
-- +migrate Up
create table if not exists product_hashtags (
    id serial not null primary KEY,
    name text,
    slug text,
    description text,
    title text,
    image text,
    color text,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

create table if not exists product_hashtag_alias (
    id serial not null primary KEY,
    product_hashtag_id integer not null references product_hashtags(id),
    alias text,
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);

insert into product_hashtags (id, name, slug, description, title, image, color) VALUES
(1, 'Happy birthday', 'happy_birthday', 'You''re not getting older, you''re just leveling up. So enjoy your age!

{.user} sent you {.token_icon}  **{.amount} {.token}** (= $${.usd_amount}) as birthday gift.
<:chat:1078633889247006790> Hope u happy #hpbd', '<a:party_popper:1095990305414709331> It''s someone''s birthday! Let''s celebrate!<a:party_popper:1095990305414709331>', 'https://cdn.discordapp.com/attachments/818751486162632710/1151360213182582834/4_Birthday.png', 'ffc957'),
(2, 'Achievement', 'achievement', 'You did it! <a:Agree:1133954133146226779> Whether it''s a personal milestone or a big win! 

{.user} sent you {.token_icon}  **{.amount} {.token}** (= ${.usd_amount})  to celebrate your achievement! <:trophy:1060414870895464478> <:trophy:1060414870895464478> <:trophy:1060414870895464478>
<:chat:1078633889247006790> Hope u happy #win', '<a:star:1093923083934502982> <a:star:1093923083934502982> <a:star:1093923083934502982> Celebrate New Achievement <a:star:1093923083934502982> <a:star:1093923083934502982> <a:star:1093923083934502982>', 'https://cdn.discordapp.com/attachments/818751486162632710/1151360212935135313/3_Women_day.png', 'ffc957'),
(3, 'Appreciation', 'Appreciation', '{.user} sent you {.token_icon}  **{.amount} {.token} (= ${.usd_amount})**  as an appreciation!
<:chat:1078633889247006790> Hope u happy #thanks', '<a:heart:1093923040859009025> <a:heart:1093923040859009025> <a:heart:1093923040859009025> Thank You Gift <a:heart:1093923040859009025> <a:heart:1093923040859009025> <a:heart:1093923040859009025>', 'https://cdn.discordapp.com/attachments/818751486162632710/1151403908191756289/9_Thanks_I_owe_you_one.png', 'ff5f5f'),
(4, 'Anniversary', 'Anniversary', '{.user} sent you {.token_icon}  **{.amount} {.token} (= ${.usd_amount})**  as an appreciation!
<:chat:1078633889247006790> Hope u happy #anni', '<a:heart:1093923040859009025> <a:heart:1093923040859009025> <a:heart:1093923040859009025> Happy Annivesary <a:heart:1093923040859009025> <a:heart:1093923040859009025> <a:heart:1093923040859009025>', 'https://cdn.discordapp.com/attachments/818751486162632710/1151360213962723389/7_Wedding.png', 'ff5f5f');

insert into product_hashtag_alias (product_hashtag_id, alias) values 
(1, 'happybirthday'),
(1, 'happy_birthday'),
(1, 'hpbd'),
(2, 'win'),
(2, 'achievement'),
(2, 'congrat'),
(2, 'congratulation'),
(4, 'anni'),
(4, 'anniversary'),
(4, 'valentine'),
(4, 'love'),
(3, 'thanks'),
(3, 'thank'),
(3, 'appreciate'),
(3, 'appreciation');
-- +migrate Down
drop table if exists product_hashtags;
drop table if exists product_hashtag_alias;