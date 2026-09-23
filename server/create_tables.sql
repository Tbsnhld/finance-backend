CREATE TABLE user (
  id UUID PRIMARY KEY,
  title TEXT,
  description TEXT
);

CREATE TABLE entry (
  id UUID PRIMARY KEY,
  createdBy UUID NOT NULL REFERENCES user (id),
  value FLOAT32 NOT NULL,
  users [UUID] NOT NULL,
)
