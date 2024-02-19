INSERT INTO users (
  id,
  username,
  email,
  password,
) VALUES (NULL, `admin`, `admin@admin.io`, `admin` )

INSERT INTO tasks (
  id,
  userid,
  FOREIGN KEY(userid) REFERENCES users (id),
  title,
  description,
  created_at,
  last_modified,
  due_at,
) VALUES ( NULL, `1`, `Title`, `Description`, NULL, NULL, NULL)
