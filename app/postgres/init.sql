-- Extension for generating uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Table of users
CREATE TABLE users (
  id UUID PRIMARY KEY,
  username VARCHAR(32) NOT NULL,
  email VARCHAR(64) NOT NULL UNIQUE,
  password VARCHAR(64) NOT NULL,
)

-- Table of tasks
CREATE TABLE tasks (
  id UUID PRIMARY KEY,
  userid UUID NOT NULL,
  FOREIGN KEY(userid) REFERENCES users (id),
  title VARCHAR(32) NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_modified TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  due_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
)

-- Populate sample user
INSERT INTO users (
  id,
  username,
  email,
  password
) VALUES ( uuid_generate_v4(), 'admin', 'admin@admin.io', 'admin' )

-- Populate sample task
INSERT INTO tasks (
  id,
  userid,
  title,
  description,
  created_at,
  last_modified,
  due_at
) VALUES ( uuid_generate_v4(), (SELECT id FROM users WHERE username = 'admin'), 'Title', 'Description', NULL, NULL, NULL)
