CREATE TABLE IF NOT EXISTS "user"(
   user_uuid char(36) PRIMARY KEY,
   user_username varchar(15) NOT NULL,
   user_password varchar(20) NOT NULL,
   user_about_me text,
   user_joined_date date NOT NULL,
   user_is_banned bool NOT NULL DEFAULT FALSE,
   user_is_admin bool NOT NULL DEFAULT FALSE
);

