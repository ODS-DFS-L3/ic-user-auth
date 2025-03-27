CREATE TABLE public.plants (
    plant_id character varying(256) NOT NULL,
    operator_id character varying(256) NOT NULL,
    plant_name character varying(256) NOT NULL,
    plant_address character varying(256) NOT NULL,
    deleted_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    created_user_id text NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    updated_user_id text NOT NULL,
    open_plant_id character varying(26) NOT NULL,
    global_plant_id character varying(256)
);

COMMENT ON TABLE public.plants IS '事業所テーブル';
COMMENT ON COLUMN public.plants.plant_id IS '事業所識別子';
COMMENT ON COLUMN public.plants.operator_id IS '事業者識別子（外部Key）';
COMMENT ON COLUMN public.plants.plant_name IS '事業所名';
COMMENT ON COLUMN public.plants.plant_address IS '事業所所在値（住所）';
COMMENT ON COLUMN public.plants.deleted_at IS '論理削除日時';
COMMENT ON COLUMN public.plants.created_at IS '作成日時';
COMMENT ON COLUMN public.plants.created_user_id IS '作成ユーザ';
COMMENT ON COLUMN public.plants.updated_at IS '更新日時';
COMMENT ON COLUMN public.plants.updated_user_id IS '更新ユーザ';
COMMENT ON COLUMN public.plants.open_plant_id IS '公開事業所識別子';
COMMENT ON COLUMN public.plants.global_plant_id IS '事業所識別子（グローバル）';

ALTER TABLE ONLY public.plants ADD CONSTRAINT plants_pkey PRIMARY KEY (plant_id, operator_id);
ALTER TABLE ONLY public.plants ADD CONSTRAINT unique_global_plant_id_operator_id UNIQUE (operator_id, global_plant_id);
ALTER TABLE ONLY public.plants ADD CONSTRAINT unique_open_plant_id_operator_id UNIQUE (operator_id, open_plant_id);
ALTER TABLE ONLY public.plants ADD CONSTRAINT plants_operator_id_fkey FOREIGN KEY (operator_id) REFERENCES public.operators(operator_id) ON UPDATE CASCADE ON DELETE CASCADE;

INSERT INTO public.plants (plant_id, operator_id, plant_name, plant_address, deleted_at, created_at, created_user_id, updated_at, updated_user_id, open_plant_id, global_plant_id) VALUES ('00000000-0000-0000-0000-000000000101','b39e6248-c888-56ca-d9d0-89de1b1adc8e', '事業所A1', '工場A1', NULL, '2024-05-01 00:00:00.000000', 'seed', '2024-05-01 00:00:00.000000', 'seed', '000000000001', '000000000001');
INSERT INTO public.plants (plant_id, operator_id, plant_name, plant_address, deleted_at, created_at, created_user_id, updated_at, updated_user_id, open_plant_id, global_plant_id) VALUES ('00000000-0000-0000-0000-000000000102','b39e6248-c888-56ca-d9d0-89de1b1adc8e', '事業所A2', '工場A2', NULL, '2024-05-01 00:00:00.000000', 'seed', '2024-05-01 00:00:00.000000', 'seed', '000000000002', '000000000002');
INSERT INTO public.plants (plant_id, operator_id, plant_name, plant_address, deleted_at, created_at, created_user_id, updated_at, updated_user_id, open_plant_id, global_plant_id) VALUES ('00000000-0000-0000-0000-000000000111','15572d1c-ec13-0d78-7f92-dd4278871373', '事業所B1', '工場B1', NULL, '2024-05-01 00:00:00.000000', 'seed', '2024-05-01 00:00:00.000000', 'seed', '000000000011', '000000000011');
INSERT INTO public.plants (plant_id, operator_id, plant_name, plant_address, deleted_at, created_at, created_user_id, updated_at, updated_user_id, open_plant_id, global_plant_id) VALUES ('00000000-0000-0000-0000-000000000112','15572d1c-ec13-0d78-7f92-dd4278871373', '事業所B2', '工場B2', NULL, '2024-05-01 00:00:00.000000', 'seed', '2024-05-01 00:00:00.000000', 'seed', '000000000012', '000000000012');