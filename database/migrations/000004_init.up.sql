CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

ALTER TABLE users
ALTER COLUMN id SET DEFAULT uuid_generate_v4();