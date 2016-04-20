create extension if not exists pgcrypto;
create extension if not exists pg_trgm;

CREATE TABLE users (
	id serial not null,
	uuid varchar(32) not null,
	created_utc timestamp not null,
	username varchar(255) not null,
	first_name varchar(64),
	last_name varchar(64),
    email_address varchar(255),
    is_email_verified boolean not null,
    is_admin boolean not null default false,
	is_banned boolean not null default false
);
ALTER TABLE users ADD CONSTRAINT pk_users_id PRIMARY KEY (id);
ALTER TABLE users ADD CONSTRAINT uk_users_uuid UNIQUE (uuid);
ALTER TABLE users ADD CONSTRAINT uk_users_username UNIQUE (username);

CREATE TABLE user_auth (
    user_id bigint not null,
    provider varchar(32) not null,
    timestamp_utc timestamp not null,
    auth_token bytea not null,
	auth_token_hash bytea not null,
    auth_secret bytea
);
ALTER TABLE user_auth ADD CONSTRAINT pk_user_auth_user_id_provider PRIMARY KEY (user_id,provider);
ALTER TABLE user_auth ADD CONSTRAINT fk_user_auth_user_id FOREIGN KEY (user_id) REFERENCES users(id);
CREATE INDEX ix_user_auth_auth_token_hash ON user_auth(auth_token_hash);

CREATE TABLE user_session (
	session_id varchar(32) not null,
    user_id bigint not null,
    timestamp_utc timestamp not null
);
ALTER TABLE user_session ADD CONSTRAINT pk_user_session_session_id PRIMARY KEY (session_id);
ALTER TABLE user_session ADD CONSTRAINT fk_user_session_user_id FOREIGN KEY (user_id) REFERENCES users(id);
