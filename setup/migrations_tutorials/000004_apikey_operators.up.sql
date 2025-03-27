CREATE TABLE public.apikey_operators (
    api_key character varying(256) NOT NULL,
    operator_id character varying(256) NOT NULL,
    deleted_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    created_user_id text NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    updated_user_id text NOT NULL
);

COMMENT ON COLUMN public.apikey_operators.api_key IS 'APIキー(外部Key)';
COMMENT ON COLUMN public.apikey_operators.operator_id IS '事業者識別子（外部Key）';
COMMENT ON COLUMN public.apikey_operators.deleted_at IS '論理削除日時';
COMMENT ON COLUMN public.apikey_operators.created_at IS '作成日時';
COMMENT ON COLUMN public.apikey_operators.created_user_id IS '作成ユーザ';
COMMENT ON COLUMN public.apikey_operators.updated_at IS '更新日時';
COMMENT ON COLUMN public.apikey_operators.updated_user_id IS '更新ユーザ';

ALTER TABLE ONLY public.apikey_operators ADD CONSTRAINT apikey_operators_pkey PRIMARY KEY (api_key, operator_id);
ALTER TABLE ONLY public.apikey_operators ADD CONSTRAINT apikey_operators_api_key_fkey FOREIGN KEY (api_key) REFERENCES public.api_keys(api_key) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE ONLY public.apikey_operators ADD CONSTRAINT apikey_operators_operator_id_fkey FOREIGN KEY (operator_id) REFERENCES public.operators(operator_id) ON UPDATE CASCADE ON DELETE CASCADE;

INSERT INTO public.apikey_operators(api_key, operator_id, deleted_at, created_at, created_user_id, updated_at, updated_user_id)VALUES('Sample-APIKey1', 'b39e6248-c888-56ca-d9d0-89de1b1adc8e',  NULL, '2024-05-01 00:00:00.000000', 'seed', '2024-05-01 00:00:00.000000', 'seed');
INSERT INTO public.apikey_operators(api_key, operator_id, deleted_at, created_at, created_user_id, updated_at, updated_user_id)VALUES('Sample-APIKey1', '15572d1c-ec13-0d78-7f92-dd4278871373',  NULL, '2024-05-01 00:00:00.000000', 'seed', '2024-05-01 00:00:00.000000', 'seed');