-- 银行信息
DROP TABLE IF EXISTS public.nice_bank_info;
DROP SEQUENCE IF EXISTS public.nice_bank_info_id_seq;

CREATE SEQUENCE nice_bank_info_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE nice_bank_info(
    id INTEGER NOT NULL DEFAULT nextval('nice_bank_info_id_seq'::regclass),
    bank_province VARCHAR(50) DEFAULT '' NOT NULL,
    bank_no VARCHAR(50) DEFAULT '' NOT NULL,
    bank_name VARCHAR(100) DEFAULT '' NOT NULL,
    mobile VARCHAR(30) DEFAULT '' NOT NULL,
    zip_code VARCHAR(30) DEFAULT '' NOT NULL,
    address text
);

COMMENT ON TABLE "public"."nice_bank_info" IS '银行信息';
COMMENT ON COLUMN "public"."nice_bank_info"."id" IS 'id';
COMMENT ON COLUMN "public"."nice_bank_info"."bank_province" IS '省份';
COMMENT ON COLUMN "public"."nice_bank_info"."bank_no" IS '行号';
COMMENT ON COLUMN "public"."nice_bank_info"."bank_name" IS '名称';
COMMENT ON COLUMN "public"."nice_bank_info"."mobile" IS '电话';
COMMENT ON COLUMN "public"."nice_bank_info"."zip_code" IS '邮编';
COMMENT ON COLUMN "public"."nice_bank_info"."address" IS '地址';

ALTER TABLE ONLY nice_bank_info ADD CONSTRAINT nice_bank_info_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX idx_bank_no ON nice_bank_info(bank_no);
