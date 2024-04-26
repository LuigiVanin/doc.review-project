CREATE TABLE participants (
	id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id),
  homework_id UUID NOT NULL REFERENCES homework(id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT user_homework_unique UNIQUE (user_id, homework_id)
);

ALTER TABLE documents
ADD COLUMN submitted_homework_id UUID REFERENCES homework(id) DEFAULT NULL;