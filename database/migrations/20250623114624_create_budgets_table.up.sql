CREATE TABLE budgets (
	id SERIAL PRIMARY KEY NOT NULL,
	user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	account_id BIGINT NOT NULL REFERENCES accounts(id) ON DELETE CASCADE,
	category_id BIGINT NOT NULL REFERENCES categories(id),
	name VARCHAR(100) NOT NULL,
	amount DECIMAL(15,2) NOT NULL,
	period_type recurring_frequency,
	start_date DATE NOT NULL,
	end_date DATE,
	alert_percentage DECIMAL(5,2) DEFAULT 80.00,
	is_active BOOLEAN DEFAULT TRUE,
	created_at timestamp WITH TIME ZONE DEFAULT NOW(),
	updated_at timestamp WITH TIME ZONE DEFAULT NOW(),
	deleted_at timestamp WITH TIME ZONE NULL
);
	