
-- +migrate Up
ALTER TABLE product_hashtags ADD column alias text[];

UPDATE product_hashtags set alias = '{happybirthday,happy_birthday,hpbd}' where id = 1;
UPDATE product_hashtags set alias = '{win,achievement,congrat,congratulation}' where id = 2;
UPDATE product_hashtags set alias = '{thanks,thank,appreciate,appreciation}' where id = 3;
UPDATE product_hashtags set alias = '{anni,anniversary,valentine,love}' where id = 4;
-- +migrate Down
ALTER  TABLE product_hashtags DROP column alias;
