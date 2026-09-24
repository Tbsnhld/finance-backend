CREATE TABLE application_user (
  id UUID PRIMARY KEY,
  title TEXT,
  description TEXT
);

CREATE TABLE tag (
  id UUID PRIMARY KEY,
  name varchar(255) NOT NULL 
);

CREATE TABLE category (
  id UUID PRIMARY KEY,
  name varchar(255) NOT NULL 
);


CREATE TABLE entry (
  id UUID PRIMARY KEY,
  createdBy UUID NOT NULL REFERENCES application_user (id),
  value REAL NOT NULL,
  users UUID[] NOT NULL,
  split REAL[],
  time DATE
);

CREATE TABLE entry_tag (
  entry_id UUID NOT NULL REFERENCES entry (id),
  tag_id UUID NOT NULL REFERENCES tag (id),
  PRIMARY KEY (entry_id, tag_id)
);

CREATE TABLE entry_category (
  entry_id UUID NOT NULL REFERENCES entry (id),
  category_id UUID NOT NULL REFERENCES category (id),
  PRIMARY KEY (entry_id, category_id)
);
