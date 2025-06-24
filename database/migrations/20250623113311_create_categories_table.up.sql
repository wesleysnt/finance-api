CREATE TABLE categories (
	id SERIAL PRIMARY KEY NOT NULL,
	user_id BIGINT NOT NULL,
	name VARCHAR(100) NOT NULL,
	category_type type,
	parent_id BIGINT REFERENCES categories(id),
	color VARCHAR(7),
	icon VARCHAR(50),
	is_active BOOLEAN DEFAULT TRUE,
	created_at timestamp WITH TIME ZONE DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE DEFAULT NOW(),
	deleted_at timestamp WITH TIME ZONE NULL
);
	