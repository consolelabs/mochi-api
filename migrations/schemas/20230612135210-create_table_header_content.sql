
-- +migrate Up
create table if not exists contents (
    id serial primary key,
    type varchar,
    description json,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

insert into contents (type, description) values ('header', 
    '{
        "tip": ["Use /feedback to report and get reward", "Check your watchlist /wlv thrice a day to get more XP", "Run /ticker thrice a day to get more XP", "Use /earn to maximize your earning", "Mochi /profile help create 10 wallets", "Use /tip to pay your frens"],
        "fact": ["Mochi もち, 餅 origin means **Japanese Rice Cake**", "Mochi (pronounced MOE-chee) is a Japanese dessert made of sweet glutinous rice flour or mochigome", "Mochi dough is often tinted with green tea powder", "In Korean, mochi is pronounced mojji 모찌"]
    }'
);
-- +migrate Down
drop table if exists contents;
