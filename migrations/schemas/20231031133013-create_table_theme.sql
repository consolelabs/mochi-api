
-- +migrate Up
create table if not exists product_themes (
    id serial primary key,
    name text,
    slug text,
    image text,
    created_at timestamp default NOW(),
    updated_at timestamp default NOW()
);

insert into product_themes (id, name, slug, image) VALUES 
    (1, 'Happy birthday', 'happy_birthday', 'https://cdn.discordapp.com/attachments/818751486162632710/1151360213182582834/4_Birthday.png'),
    (2, 'Achievement', 'achievement', 'https://cdn.discordapp.com/attachments/818751486162632710/1151360212935135313/3_Women_day.png'),
    (3, 'Appreciation', 'appreciation', 'https://cdn.discordapp.com/attachments/818751486162632710/1151403908191756289/9_Thanks_I_owe_you_one.png'),
    (4, 'Anniversary', 'anniversary', 'https://cdn.discordapp.com/attachments/818751486162632710/1151360213962723389/7_Wedding.png');
-- +migrate Down
drop table if exists product_themes;
