CREATE TYPE user_type AS ENUM ('Professor', 'Aluno');

CREATE TABLE users (
	Id UUID NOT NULL PRIMARY KEY,
	Name TEXT NOT NULL,
	Email TEXT NOT NULL UNIQUE,
	Type user_type NOT NULL,
	Password TEXT NOT NULL
);