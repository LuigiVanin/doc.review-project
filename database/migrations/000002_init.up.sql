-- TODO: add updated_at and deadline columns to the homework table

CREATE TABLE homework (
	id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
	title VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
  owner_id UUID NOT NULL REFERENCES users(id),
  deleted boolean NOT NULL DEFAULT FALSE,
  code TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT code_unique UNIQUE (code)
);