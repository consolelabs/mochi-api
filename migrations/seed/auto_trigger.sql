
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
INSERT INTO public.auto_triggers(id, discord_guild_id, user_discord_id, name, status, updated_at, created_at) VALUES (1, '711823851117608990', '711823390000021556', 'Trigger when send Message', true, now(), now());
INSERT INTO public.auto_triggers(id, discord_guild_id, user_discord_id, name, status, updated_at, created_at) VALUES (2, '711823851117608990', '711823390000021556', 'React heart in general', true, now(), now());

-- Seed data
INSERT INTO public.auto_actions(id, trigger_id, type_id, channel_ids, name, content, action_data, created_at, index) VALUES (1, 1, 2, '', 'Do action 1', 'Hello world', '', now(), 1);
INSERT INTO public.auto_actions(id, trigger_id, type_id, channel_ids, name, content, action_data, created_at, index) VALUES (2, 2, 2, '', 'Do action 2', 'Hello world', '', now(), 2);
INSERT INTO public.auto_actions(id,  trigger_id, type_id, channel_ids, name, content, action_data, created_at, index) VALUES (3,2, 8, '', 'Do action 3', 'Hello world 3', '', now(), 3);

-- Seed data condition
INSERT INTO public.auto_conditions(id, trigger_id, type_id, channel_id, index, platform, updated_at, created_at) VALUES (1, 2, 1, '711823851117608992', 1, 'discord', now(), now());

-- Seed data condition values
INSERT INTO public.auto_condition_values(id, condition_id, type_id, index, operator, created_at, matches) VALUES (1, 1, 7, 1, '', now(), '');
INSERT INTO public.auto_condition_values(id, condition_id, type_id, index, operator, created_at, matches) VALUES (2, 1, 7, 2, 'in', now(), '');
INSERT INTO public.auto_condition_values(id, condition_id, type_id, index, operator, created_at, matches) VALUES (3, 1, 7, 1, 'in', now(), '1115906135799648257,711823851117608990');
INSERT INTO public.auto_condition_values(id, condition_id, type_id, index, operator, created_at, matches) VALUES (4, 1, 9, 2, '==', now(), '<:pepeheart:867454854686048256>');

INSERT INTO public.auto_embeds(id, author_id, action_id, title, description, color, url, type, fields, created_at)
	VALUES (1, 'fbaca19d-4ecc-4627-92b6-81858536f921', 2, 'Embed 1', 'descript 1', '#000000', 'https://openai.com/research/measuring-goodharts-law', 'some types', 'any fields', now());

INSERT INTO public.auto_embed_images(id, embed_id, url, proxy_url, height, width, created_at)
	VALUES ('fb64b906-7e0c-4162-b79d-2f0690b543da', 1, 'https://cafefcdn.com/thumb_w/640/203337114487263232/2023/6/15/avatar1686801739290-16868017399091383965775.jpg', 'https://cafefcdn.com/thumb_w/640/203337114487263232/2023/6/15/avatar1686801739290-16868017399091383965775.jpg', 100, 100, now());

INSERT INTO public.auto_embed_footers(id, embed_id, text, icon_url, url, created_at)
	VALUES (1, 1, 'Hello footer 1', 'icon', 'https://www.investopedia.com/terms/d/deltahedging.asp', now());

INSERT INTO public.auto_embed_videos(id, embed_id, url, height, width, created_at)
	VALUES (1, 1, 'https://arxiv.org/pdf/2210.10760.pdf', 110, 50, now());


-- Seed data to action transfer vault
INSERT INTO public.auto_actions(id,  trigger_id, type_id, channel_ids, name, content, action_data, created_at) VALUES (1,2, 10, '', 'Transfer vault 1', 'Transfer vault by address', '{
    "discord_guild_id": "711823851117608990",
    "vault_id": 2,
    "message": "Send money to treasurer",
    "chain": "137",
    "token": "matic",
    "amount": "0.00001"
}', now());
