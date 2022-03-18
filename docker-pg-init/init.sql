CREATE ROLE db_admin LOGIN PASSWORD 'admin_password' NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;
CREATE ROLE db_user LOGIN PASSWORD 'user_password' NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;
CREATE ROLE db_ramdonly LOGIN PASSWORD 'ramdonly_password' NOSUPERUSER INHERIT NOCREATEDB NOCREATEROLE NOREPLICATION;

CREATE DATABASE demo WITH ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8' CONNECTION LIMIT = -1 template=template0;
\connect demo;
REVOKE USAGE ON SCHEMA public FROM PUBLIC;
REVOKE CREATE ON SCHEMA public FROM PUBLIC;

GRANT USAGE ON SCHEMA public TO db_admin;
GRANT CREATE ON SCHEMA public TO db_admin;

/* grant the schema access privilege to normal users. Without schema right, user will unable to see the tables. */
GRANT USAGE ON SCHEMA public TO db_user;
GRANT USAGE ON SCHEMA public to db_ramdonly;


CREATE TABLE user_table 
(
	username VARCHAR(64) NOT NULL,
	passwd VARCHAR(64) NOT NULL,
	CONSTRAINT "user_name" PRIMARY KEY (username)
);

ALTER TABLE user_table OWNER TO db_admin;

/* grant the CRUD privilege to normal users. grant the select privilege to ramdon users. */
GRANT SELECT, INSERT, UPDATE, DELETE, REFERENCES ON TABLE user_table TO db_user;
GRANT SELECT ON TABLE user_table TO db_ramdonly;