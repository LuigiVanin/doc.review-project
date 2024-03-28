CREATE TYPE user_type AS ENUM ('Professor', 'Aluno');

CREATE TABLE users (
	id UUID NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	type user_type NOT NULL,
	password TEXT NOT NULL
);