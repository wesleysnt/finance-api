CREATE TABLE transactions (
	id SERIAL PRIMARY KEY NOT NULL,
	user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
	category_id BIGINT NOT NULL REFERENCES categories(id),
	amount DECIMAL(15,2) NOT NULL,
	description TEXT,
	transaction_date Date NOT NULL,
	transaction_type type,
	referece_number VARCHAR(20),
	recurring_frequency recurring_frequency,
	is_opening_balance BOOLEAN DEFAULT FALSE,
	tags TEXT[],
	created_at timestamp WITH TIME ZONE DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE DEFAULT NOW(),
	deleted_at timestamp WITH TIME ZONE NULL
);
