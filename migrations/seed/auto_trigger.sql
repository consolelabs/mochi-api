
-- Seed data to auto_types
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES ('2236b7a8-9f9c-456b-a5fa-2dc0755d24b1', 'create message', 'createMessage', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES ('2236b7a8-9f9c-456b-a5fa-2dc0755d24b2', 'React message', 'reactionAdd', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES ('2236b7a8-9f9c-456b-a5fa-2dc0755d24b3', 'Total react', 'totalReact', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES ('2236b7a8-9f9c-456b-a5fa-2dc0755d24b4', 'Total message', 'totalMessage', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES ('2236b7a8-9f9c-456b-a5fa-2dc0755d24b5', 'Message in channel', 'messageChannel', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES ('2236b7a8-9f9c-456b-a5fa-2dc0755d24b6', 'React in channel', 'reactChannel', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES ('2236b7a8-9f9c-456b-a5fa-2dc0755d24b7', 'user role', 'userRole', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES ('2236b7a8-9f9c-456b-a5fa-2dc0755d24b8', 'action send message', 'sendMessage', '', now());
INSERT INTO public.auto_types(id, name, type, icon_url, created_at) VALUES ('2236b7a8-9f9c-456b-a5fa-2dc0755d24b9', 'React type', 'reactType', '', now());

-- Seed data trigger
INSERT INTO public.auto_triggers(id, guild_id, user_id, name, status, updated_at, created_at) VALUES ('7236b7a8-9f9c-456b-a5fa-2dc0755d24bb', '711823851117608990', '711823390000021556', 'Trigger when send Message', true, now(), now());
INSERT INTO public.auto_triggers(id, guild_id, user_id, name, status, updated_at, created_at) VALUES ('7236b7a8-9f9c-456b-a5fa-2dc0755d24ba', '711823851117608990', '711823390000021556', 'React heart in general', true, now(), now());

-- Seed data 
INSERT INTO public.auto_actions(id, trigger_id, type_id, channel_ids, name, content, action_data, created_at) VALUES ('4236b7a8-9f9c-456b-a5fa-2dc0755d24b2', '7236b7a8-9f9c-456b-a5fa-2dc0755d24bb', '2236b7a8-9f9c-456b-a5fa-2dc0755d24b2', '', 'Do action 1', 'Hello world', '', now());
INSERT INTO public.auto_actions(id, trigger_id, type_id, channel_ids, name, content, action_data, created_at) VALUES ('4236b7a8-9f9c-456b-a5fa-2dc0755d24b3', '7236b7a8-9f9c-456b-a5fa-2dc0755d24ba', '2236b7a8-9f9c-456b-a5fa-2dc0755d24b2', '', 'Do action 2', 'Hello world', '', now());
INSERT INTO public.auto_actions(id,  trigger_id, type_id, channel_ids, name, content, action_data, created_at) VALUES ('4236b7a8-9f9c-456b-a5fa-2dc0755d24b4','7236b7a8-9f9c-456b-a5fa-2dc0755d24ba', '2236b7a8-9f9c-456b-a5fa-2dc0755d24b8', '', 'Do action 3', 'Hello world 3', '', now());

-- Seed data condition
INSERT INTO public.auto_conditions(id, trigger_id, type_id, channel_id, index, platform, updated_at, created_at) VALUES ('7236b7a8-9f9c-456b-a5fa-2dc0755d24b3', '7236b7a8-9f9c-456b-a5fa-2dc0755d24ba', '7236b7a8-9f9c-456b-a5fa-2dc0755d24b2', '711823851117608992', 1, 'discord', now(), now());

-- Seed data condition values
INSERT INTO public.auto_condition_values(id, condition_id, type_id, index, operator, created_at, matches) VALUES ('7236b7a8-9f9c-456b-a5fa-2dc0755d24b1', '7236b7a8-9f9c-456b-a5fa-2dc0755d24b2', '2236b7a8-9f9c-456b-a5fa-2dc0755d24b7', 1, '', now(), '');
INSERT INTO public.auto_condition_values(id, condition_id, type_id, index, operator, created_at, matches) VALUES ('7236b7a8-9f9c-456b-a5fa-2dc0755d24b2', '7236b7a8-9f9c-456b-a5fa-2dc0755d24b2', '2236b7a8-9f9c-456b-a5fa-2dc0755d24b7', 2, 'in', now(), '');
INSERT INTO public.auto_condition_values(id, condition_id, type_id, index, operator, created_at, matches) VALUES ('7236b7a8-9f9c-456b-a5fa-2dc0755d24b5', '7236b7a8-9f9c-456b-a5fa-2dc0755d24b3', '2236b7a8-9f9c-456b-a5fa-2dc0755d24b7', 1, 'in', now(), '1115906135799648257,711823851117608990');
INSERT INTO public.auto_condition_values(id, condition_id, type_id, index, operator, created_at, matches) VALUES ('7236b7a8-9f9c-456b-a5fa-2dc0755d24b7', '7236b7a8-9f9c-456b-a5fa-2dc0755d24b3', '2236b7a8-9f9c-456b-a5fa-2dc0755d24b9', 2, '==', now(), "<:pepeheart:867454854686048256>");
INSERT INTO public.auto_condition_values(id, condition_id, type_id, index, operator, created_at, matches) VALUES ('7236b7a8-9f9c-456b-a5fa-2dc0755d24b6', '7236b7a8-9f9c-456b-a5fa-2dc0755d24b3', '2236b7a8-9f9c-456b-a5fa-2dc0755d24b3', 3, '==', now(), "1");

INSERT INTO public.auto_embeds(id, author_id, action_id, title, description, color, url, type, fields, created_at)
	VALUES ('d95d4370-2629-4d85-85ea-b8d623dbaff4', 'fbaca19d-4ecc-4627-92b6-81858536f921', '4236b7a8-9f9c-456b-a5fa-2dc0755d24b3', 'Embed 1', 'descript 1', '#000000', 'https://openai.com/research/measuring-goodharts-law', 'some types', 'any fields', now(), now());

INSERT INTO public.auto_embed_images(id, embed_id, url, proxy_url, height, width, created_at)
	VALUES ('fb64b906-7e0c-4162-b79d-2f0690b543da', 'd95d4370-2629-4d85-85ea-b8d623dbaff4', 'https://cafefcdn.com/thumb_w/640/203337114487263232/2023/6/15/avatar1686801739290-16868017399091383965775.jpg', 'https://cafefcdn.com/thumb_w/640/203337114487263232/2023/6/15/avatar1686801739290-16868017399091383965775.jpg', 100, 100, now());

INSERT INTO public.auto_embed_footers(id, embed_id, text, icon_url, url, created_at)
	VALUES ('e0a1870a-8cfc-407d-9492-e8a45c0d613d', 'd95d4370-2629-4d85-85ea-b8d623dbaff4', 'Hello footer 1', 'icon', 'https://www.investopedia.com/terms/d/deltahedging.asp', now());

INSERT INTO public.auto_embed_videos(id, embed_id, url, height, width, created_at)
	VALUES (uuid_generate_v4(), 'd95d4370-2629-4d85-85ea-b8d623dbaff4', 'https://arxiv.org/pdf/2210.10760.pdf', 110, 50, now());