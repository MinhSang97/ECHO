-- +migrate Up
CREATE table "users"
(
    "user_id"    text primary key,
    "full_name"  text,
    "email"      text unique,
    "password"   text,
    "role"       text,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" timestamp not NULL

);

CREATE table "repos"
(
    "name"        text primary key,
    "description" text,
    "url"         text,
    "color"       text,
    "lang"        text,
    "fork"        text,
    "stars"       text,
    "stars_today" text,
    "build_by"    text,
    "created_at"  TIMESTAMP not null,
    "updated_at"  Timestamp not null
);

CREATE TABLE "bookmarks"
(
    "bid"        text PRIMARY KEY,
    "user_id"    text,
    "repo_name"  text UNIQUE,
    "created_at" TIMESTAMPTZ NOT NULL,
    "updated_at" TIMESTAMPTZ NOT NULL
);


ALTER TABLE "bookmarks"
    ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");
ALTER TABLE "bookmarks"
    ADD FOREIGN KEY ("repo_name") REFERENCES "repos" ("name");

-- +migrate Down
DROP TABLE users;
DROP TABLE repos;
DROP TABLE bookmarks;