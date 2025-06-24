CREATE TABLE users (
	id SERIAL PRIMARY KEY NOT NULL,
	name VARCHAR NOT NULL,
	email VARCHAR NOT NULL UNIQUE,
	password varchar NOT NULL,
	currency VARCHAR(3) DEFAULT 'IDR',
	created_at timestamp WITH TIME ZONE DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE DEFAULT NOW(),
	deleted_at timestamp WITH TIME ZONE NULL
);

CREATE INDEX idx_users_email on users(email);