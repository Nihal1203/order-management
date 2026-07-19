CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);

CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  firstName VARCHAR(250),
  lastName VARCHAR(250),
  email VARCHAR(250),
  password VARCHAR(250),
  line1 VARCHAR(250),
  line2 VARCHAR(250),
  city VARCHAR(250),
  state VARCHAR(250),
  zipcode INT);

CREATE TABLE orders (
  id   BIGSERIAL PRIMARY KEY,
  userid BIGINT,
  itemName VARCHAR(250),
  quantity BIGINT,
  status VARCHAR(250),
  instruction TEXT DEFAULT '',
  CONSTRAINT fk_users FOREIGN KEY (userid) REFERENCES users(id)
);  