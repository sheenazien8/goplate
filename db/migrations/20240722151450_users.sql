-- migrate:up
CREATE SEQUENCE users_id_seq START WITH 1;
CREATE TABLE "users"(
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "username" TEXT,
    "password" TEXT,
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    PRIMARY KEY("id"),
    UNIQUE("username")
);

-- migrate:down
DROP TABLE "users";
DROP SEQUENCE users_id_seq;
