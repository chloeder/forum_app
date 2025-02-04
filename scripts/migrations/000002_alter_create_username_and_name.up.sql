ALTER TABLE users
ADD username VARCHAR(255) NOT NULL;

ALTER TABLE users
ADD name VARCHAR(255) NOT NULL;

ALTER TABLE users
ADD CONSTRAINT users_username_unique UNIQUE (username);
