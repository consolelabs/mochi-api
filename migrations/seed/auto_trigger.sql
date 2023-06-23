
-- Seed data to auto_types
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (1, 'create message', 'createMessage', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (2, 'React message', 'reactionAdd', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (3, 'Total react', 'totalReact', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (4, 'Total message', 'totalMessage', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (5, 'Message in channel', 'messageChannel', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (6, 'React in channel', 'reactChannel', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (7, 'user role', 'userRole', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (8, 'action send message', 'sendMessage', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (9, 'React type', 'reactType', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES (10, 'Action transfer vault', 'vaultTransfer', '', now());

-- Seed data trigger
INSERT INTO public.auto_triggers(id, discord_guild_id, user_discord_id, name, status, updated_at, created_at) VALUES (1, '462663954813157376', '567326528216760320', 'Trigger when send Message', true, now(), now());
INSERT INTO public.auto_triggers(id, discord_guild_id, user_discord_id, name, status, updated_at, created_at) VALUES (2, '462663954813157376', '567326528216760320', 'React heart in general', true, now(), now());

-- Seed data
INSERT INTO "public"."auto_actions" ("id", "user_ids", "trigger_id", "type_id", "channel_ids", "index", "action_data", "name", "content", "then_action_id", "limit_per_user", "created_at") VALUES
(1, NULL, 1, 2, '', 1, '', 'Do action 1', 'Hello world', NULL, 1, '2023-06-23 05:57:07.797197'),
(2, NULL, 2, 2, '', 2, '', 'Do action 2', 'Hello world', NULL, 1, '2023-06-23 05:57:09.539787'),
(3, NULL, 2, 8, '', 3, '', 'Do action 3', 'Hello world 3', NULL, 1, '2023-06-23 05:57:10.905286');


-- Seed data condition
INSERT INTO public.auto_conditions(id, trigger_id, type_id, channel_id, index, platform, updated_at, created_at) VALUES (1, 2, 3, '1072722777687199744', 1, 'discord', now(), now());

-- Seed data condition values
INSERT INTO "public"."auto_condition_values" ("id", "condition_id", "child_id", "type_id", "index", "operator", "matches", "created_at") VALUES
(3, 1, NULL, 7, 1, 'in', '1115906135799648257,711823851117608990,820988147399393322,462663954813157376', '2023-06-23 06:00:13.571917'),
(4, 1, NULL, 3, 2, '==', '<:pepeheart:867454854686048256>', '2023-06-23 06:00:33.452028');

INSERT INTO public.auto_embeds(id, author_id, action_id, title, description, color, url, type, fields, created_at)
	VALUES (1, 'fbaca19d-4ecc-4627-92b6-81858536f921', 2, 'Embed 1', 'descript 1', '#000000', 'https://openai.com/research/measuring-goodharts-law', 'some types', 'any fields', now());

INSERT INTO public.auto_embed_images(id, embed_id, url, proxy_url, height, width, created_at)
	VALUES ('fb64b906-7e0c-4162-b79d-2f0690b543da', 1, 'https://cafefcdn.com/thumb_w/640/203337114487263232/2023/6/15/avatar1686801739290-16868017399091383965775.jpg', 'https://cafefcdn.com/thumb_w/640/203337114487263232/2023/6/15/avatar1686801739290-16868017399091383965775.jpg', 100, 100, now());

INSERT INTO public.auto_embed_footers(id, embed_id, text, icon_url, url, created_at)
	VALUES (1, 1, 'Hello footer 1', 'icon', 'https://www.investopedia.com/terms/d/deltahedging.asp', now());

INSERT INTO public.auto_embed_videos(id, embed_id, url, height, width, created_at)
	VALUES (1, 1, 'https://arxiv.org/pdf/2210.10760.pdf', 110, 50, now());


-- Seed data to action transfer vault
INSERT INTO public.auto_actions(id,  trigger_id, type_id, channel_ids, name, content, action_data, created_at, index) VALUES (4, 2, 10, '', 'Transfer vault 1', 'Transfer vault by address', '{
    "discord_guild_id": "711823851117608990",
    "vault_id": 2,
    "message": "Send money to treasurer",
    "chain": "137",
    "token": "matic",
    "amount": "0.00001"
}', now(), 4);
