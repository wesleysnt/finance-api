CREATE TABLE financial_goals (
	id SERIAL PRIMARY KEY NOT NULL,
	user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
	name VARCHAR(100) NOT NULL,
	description TEXT,
	target_amount DECIMAL(15,2) NOT NULL,
	current_amount DECIMAL(15,2) NOT NULL DEFAULT 0.00,
	target_date DATE,
	goal_type VARCHAR(20) NOT NULL, 
	priority INTEGER NOT NULL DEFAULT 1,
	is_achieved BOOLEAN DEFAULT FALSE,
	achieved_date DATE,
	created_at timestamp WITH TIME ZONE DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE DEFAULT NOW(),
	deleted_at timestamp WITH TIME ZONE NULL
	);
	