/*
 Navicat Premium Data Transfer

 Source Server         : local-postgres
 Source Server Type    : PostgreSQL
 Source Server Version : 140001
 Source Host           : localhost:5432
 Source Catalog        : sawitpro
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140001
 File Encoding         : 65001

 Date: 14/02/2024 22:20:26
*/


-- ----------------------------
-- Sequence structure for user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."user_id_seq";
CREATE SEQUENCE "public"."user_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
                                  "id" int4 NOT NULL GENERATED ALWAYS AS IDENTITY (
                                      INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
),
                                  "full_name" varchar(255) COLLATE "pg_catalog"."default",
                                  "phone_number" varchar(255) COLLATE "pg_catalog"."default",
                                  "password" varchar(255) COLLATE "pg_catalog"."default",
                                  "login_attempt" int4 DEFAULT 0
)
;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."user_id_seq"
    OWNED BY "public"."users"."id";
SELECT setval('"public"."user_id_seq"', 28, true);

-- ----------------------------
-- Uniques structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "phone_number" UNIQUE ("phone_number");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "public"."users" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");
